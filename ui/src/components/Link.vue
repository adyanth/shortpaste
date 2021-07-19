<template>
  <v-app>
    <v-form
      ref="link"
      v-model="valid"
      class="tab pt-10"
      @submit.prevent="submit()"
    >
      <v-col>
        <v-row>
          <v-text-field
            label="ID"
            autocomplete="off"
            v-model="id"
            :rules="idRules"
            hint="Leave empty to auto generate"
            persistent-hint
            hide-details="auto"
          ></v-text-field>
        </v-row>
        <v-row class="shrink">
          <v-text-field
            label="Link"
            v-model="link"
            hint="Link to redirect to"
            persistent-hint
            :rules="linkRules"
          ></v-text-field>
        </v-row>
        <v-row>
          <v-btn
            :right="true"
            :absolute="true"
            :disabled="!valid"
            color="success"
            type="submit"
          >
            Create
          </v-btn>
        </v-row>
      </v-col>
    </v-form>
    <v-card class="mt-12">
      <v-card-title>
        <v-text-field
          v-model="search"
          append-icon="mdi-magnify"
          label="Search"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
      <v-data-table
        :headers="headers"
        :items="items"
        :items-per-page="5"
        :search="search"
        class="elevation-1"
        ><!-- eslint-disable-next-line -->
        <template #item.id="{ item }">
          <a target="_blank" :href="getLink(item.id, true)">
            {{ item.id }}
          </a>
        </template></v-data-table
      >
    </v-card>
    <v-overlay :value="popup" :absolute="false" :opacity="0.9">
      <OnLinkSuccess
        :output="output"
        @alert="
          alert = $event;
          snackbar = true;
        "
        @close="reset"
      />
    </v-overlay>
    <v-snackbar v-model="snackbar" :timeout="10000">
      {{ alert }}

      <template v-slot:action="{ attrs }">
        <v-btn
          color="pink"
          text
          v-bind="attrs"
          @click="
            alert = '';
            snackbar = false;
          "
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script>
import OnLinkSuccess from "./OnSuccess.vue";

export default {
  components: {
    OnLinkSuccess,
  },
  data: () => ({
    valid: false,
    id: "",
    link: "",
    search: "",
    headers: [
      {
        text: "ID",
        align: "start",
        value: "id",
      },
      {
        text: "Link",
        value: "link",
      },
      {
        text: "Created At",
        value: "CreatedAt",
        filterable: false,
      },
    ],
    items: [],
    popup: false,
    output: "",
    alert: "",
    snackbar: false,
    idRules: [
      (value) =>
        !value ||
        (value.length >= 3 && value.length <= 32) ||
        "ID should be 3 to 32 characters long",
    ],
    linkRules: [
      (value) => !!value || "Required.",
      (value) => isValidHttpUrl(value) || "Enter a valid URL!",
    ],
  }),
  mounted() {
    setTimeout(this.updateTable, 30000);
    this.updateTable();
  },
  methods: {
    async submit() {
      this.generateRandom();
      console.log(this.id, this.link);
      // Call API
      fetch(this.getLink(this.id), {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: this.id,
          link: this.link,
        }),
      })
        .then((res) => {
          if (!res.ok) {
            throw new Error(
              "response code: " + res.status + ", Incorrect response received"
            );
          }
          return res.json();
        })
        .then((data) => {
          if (!data.error) {
            this.output = this.getLink(this.id, true);
            this.popup = true;
            this.alert = "Link generated!";
          } else {
            this.alert =
              "Link generation failed with error: " +
              data.error +
              " message: " +
              data.message;
          }
          this.updateTable();
        })
        .catch((err) => {
          this.alert = "Link generation failed with error: " + err;
        })
        .finally(() => {
          this.snackbar = true;
        });
    },
    generateRandom() {
      if (!this.id) {
        var uuid = require("uuid");
        this.id = uuid.v4().split("-")[0];
      }
    },
    reset() {
      this.id = "";
      this.$refs.link.reset();
      this.popup = false;
    },
    updateTable() {
      fetch(this.getLink())
        .then((resp) => resp.json())
        .then((data) => {
          this.items = data["links"];
        });
    },
    getLink(id, noapi) {
      return (
        window.location.origin +
        (noapi ? "" : "/api/v1") +
        "/l/" +
        (id ? id : "")
      );
    },
  },
};

function isValidHttpUrl(string) {
  let url;

  try {
    url = new URL(string);
  } catch (_) {
    return false;
  }

  return url.protocol === "http:" || url.protocol === "https:";
}
</script>

<style>
.tab {
  width: 50vw;
}

/* .v-input {
  flex: none;
} */
</style>