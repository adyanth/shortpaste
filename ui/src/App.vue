<template>
  <v-app id="app">
    <h1 class="title">Short{Paste}</h1>
    <div class="tab-wrapper">
      <Tabs class="tabs">
        <TabItem name="Link">
          <Link />
        </TabItem>
        <TabItem name="Text">
          <Paste />
        </TabItem>
        <TabItem name="File">
          <File />
        </TabItem>
      </Tabs>
    </div>
  </v-app>
</template>

<script>
import Link from "./components/Link.vue";
import Paste from "./components/Text.vue";
import File from "./components/File.vue";

export default {
  name: "App",
  components: {
    Link,
    Paste,
    File,
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
}

.title {
  padding: 2em;
}

.tabs {
  justify-content: center;
  align-items: center;
  /* width: 80%; */
}
</style>
