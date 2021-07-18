<template>
  <div>
    <v-theme-provider root>
      <v-card class="mx-auto" width="50vw" elevation="2" outlined>
        <v-card-title class="justify-center">Link Generated!</v-card-title>
        <v-card-subtitle>Copy the link below</v-card-subtitle>
        <v-card-text>
          <v-textarea :value="output" no-resize rows="1"></v-textarea
        ></v-card-text>
        <v-card-actions class="justify-center">
          <v-btn
            @click="
              (copy() &&
                $emit('alert', 'Copied successfully!') &&
                $emit('close')) ||
                $emit('alert', 'Copy failed!')
            "
            outlined
            rounded
            text
          >
            Copy
          </v-btn>
          <v-btn @click="$emit('close')" outlined rounded text> Close </v-btn>
        </v-card-actions>
      </v-card>
    </v-theme-provider>
  </div>
  <!-- <v-list-item three-line>
      <v-list-item-content>
        <div class="text-overline mb-4">OVERLINE</div>
        <v-list-item-title class="text-h5 mb-1"> Headline 5 </v-list-item-title>
        <v-list-item-subtitle
          >Greyhound divisely hello coldly fonwderfully</v-list-item-subtitle
        >
      </v-list-item-content>
    </v-list-item>

  </v-card> -->

  <!-- <v-col>
    <v-row> Here is the link! </v-row>
    <v-row>
      {{ output }}
    </v-row>
    <v-row>
      <v-btn color="success" @click="$emit('close')"> Done </v-btn>
    </v-row>
  </v-col> -->
</template>

<script>
export default {
  props: {
    output: {
      type: String,
      default: "",
    },
  },
  methods: {
    copy() {
      if (
        document.queryCommandSupported &&
        document.queryCommandSupported("copy")
      ) {
        var textarea = document.createElement("textarea");
        textarea.textContent = this.output;
        textarea.style.position = "fixed"; // Prevent scrolling to bottom of page in Microsoft Edge.
        document.body.appendChild(textarea);
        textarea.select();
        try {
          return document.execCommand("copy"); // Security exception may be thrown by some browsers.
        } catch (ex) {
          console.warn("Copy to clipboard failed.", ex);
          return false;
        } finally {
          document.body.removeChild(textarea);
        }
      }
    },
  },
};
</script>

<style>
</style>