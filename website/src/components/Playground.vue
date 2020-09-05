<template>
  <div v-if="!isPlaygroundFinished">
    <h1>Playground</h1>
    <v-spacer class="mb-5" />
    <div v-if="!isGameActive">
      <app-gameform @selected="displayGame" />
    </div>
    <div v-if="isGameActive && players.length > 0">
      <h2>Round {{ round }}</h2>
      <v-simple-table class="players-table">
        <template v-slot:default>
          <thead>
            <tr>
              <th class="text-left">Name</th>
              <th class="text-left">Points</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="player in players" :key="player.name">
              <td class="text-left">{{ player.name }}</td>
              <td class="text-left">{{ player.points }}</td>
            </tr>
          </tbody>
        </template>
      </v-simple-table>
    </div>
    <div v-if="isGameActive && players.length === 0">
      <h2>Waiting players</h2>
      <div v-if="winner.name">
        <h3>
          Winner of last game {{ winner.name }} with points {{ winner.points }}
        </h3>
      </div>
    </div>
  </div>
</template>

<script>
import GameForm from "./GameForm.vue";

export default {
  name: "Playground",
  components: {
    "app-gameform": GameForm
  },
  data() {
    return {
      isGameActive: false
    };
  },
  computed: {
    isGameFinished: function() {
      return this.$store.getters.game.finished;
    },
    isPlaygroundFinished: function() {
      return this.$store.getters.game.finishedPlayground;
    },
    players: function() {
      return this.$store.getters.game.players;
    },
    round: function() {
      return this.$store.getters.game.round;
    },
    winner: function() {
      return this.$store.getters.game.winner;
    }
  },
  methods: {
    displayGame() {
      this.isGameActive = true;
    }
  }
};
</script>

<style></style>
