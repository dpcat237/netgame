<template>
  <div>
    <h1>Leaderboard</h1>
    <v-simple-table v-if="winners.length > 0" class="players-table">
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Points</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(player, index) in winners" :key="index">
            <td class="text-left">{{ player.name }}</td>
            <td class="text-left">{{ player.points }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <h3 v-if="winners.length === 0">No played games yet</h3>
  </div>
</template>

<script>
export default {
  name: "Leaderboard",
  computed: {
    winners: function() {
      return this.$store.getters.game.winners;
    }
  },
  mounted() {
    this.$store.dispatch("leaderboard");
  }
};
</script>

<style></style>
