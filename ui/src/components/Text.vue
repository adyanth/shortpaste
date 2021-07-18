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
          <v-textarea
            label="Text"
            v-model="text"
            hint="Text to save."
            persistent-hint
            :rules="textRules"
          ></v-textarea>
        </v-row>
        <v-row>
          <v-checkbox
            v-model="nohighlight"
            label="Disable Syntax Highlighting"
          ></v-checkbox>
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
    text: "",
    nohighlight: false,
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
    textRules: [(value) => !!value || "Required."],
  }),
  methods: {
    async submit() {
      this.generateRandom();
      console.log(this.id, this.link);
      // Call API
      fetch("/api/v1/t/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          id: this.id,
          text: this.text,
          nohighlight: this.nohighlight,
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
            this.output = window.location.origin + "/t/" + this.id;
            this.popup = true;
            this.alert = "Link generated!";
          } else {
            this.alert =
              "Link generation failed with error: " +
              data.error +
              " message: " +
              data.message;
          }
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
  },
};
</script>

<style>
.tab {
  width: 50vw;
}

/* .v-input {
  flex: none;
} */
</style>