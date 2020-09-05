<template>
  <v-app id="app" class="v-application">
    <app-leaderboard class="mb-10" />
    <app-playground />
  </v-app>
</template>

<script>
import Leaderboard from "./components/Leaderboard.vue";
import Playground from "./components/Playground.vue";

export default {
  name: "App",
  components: {
    "app-leaderboard": Leaderboard,
    "app-playground": Playground
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
  display: inherit;
}
</style>
