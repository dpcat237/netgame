import axios from "axios";
import keys from "@/constants/keys";

class Game {
  constructor() {
    this.setDefaults();
  }

  setDefaults() {
    this.name = "";
    this.errorMessage = "";
    this.finished = false;
    this.finishedPlayground = false;
    this.players = [];
    this.round = 0;
    this.signedRequestFinished = false;
    this.winner = {};
    this.winners = [];
  }
}

export default {
  actions: {
    leaderboard({ commit }) {
      axios
        .get(keys.api.domain + "v1/game/leaderboard")
        .then(function(response) {
          if (
            response.data != null &&
            response.data.winners != null &&
            response.data.winners.length > 0
          ) {
            commit("updateData", {
              key: "winners",
              value: response.data.winners
            });
          }
        })
        .catch(function(error) {
          console.log("Error fetching leaderboard data" + error);
        });
    },
    observe({ commit, dispatch }) {
      let wbs = new WebSocket(keys.websocket.url);

      wbs.onmessage = event => {
        let data = JSON.parse(event.data);
        if (data.finished) {
          commit("updateData", { key: "winner", value: data.winner });
          commit("updateData", { key: "players", value: [] });
          dispatch("leaderboard");
        } else {
          commit("updateData", { key: "players", value: data.players });
          commit("updateData", { key: "winner", value: {} });
        }
        commit("updateData", { key: "finished", value: data.finished });
        commit("updateData", {
          key: "finishedPlayground",
          value: data["finished_playground"]
        });
        commit("updateData", { key: "round", value: data.round });
      };

      wbs.onerror = error => {
        console.log("Websocket closed " + error);
        wbs.close();
      };
    },
    sign({ commit }, data) {
      commit("updateData", { key: "signedRequestFinished", value: false });
      commit("updateData", { key: "errorMessage", value: "" });
      const reqData = {
        name: data.name,
        number_one: data.numberOne,
        number_two: data.numberTwo
      };
      const options = {};

      axios
        .post(keys.api.domain + "v1/game/sign", reqData, options)
        .then(response => {
          if (response.status === 202) {
            commit("updateData", { key: "name", value: data.name });
          }
          commit("updateData", { key: "signedRequestFinished", value: true });
        })
        .catch(error => {
          if (error.response != null) {
            commit("updateData", {
              key: "errorMessage",
              value: error.response.data.message
            });
            commit("updateData", { key: "signedRequestFinished", value: true });
            return;
          }
          commit("updateData", { key: "signedRequestFinished", value: true });
          console.log("Error signing player " + error);
        });
    }
  },
  mutations: {
    updateData(state, data) {
      state[data.key] = data.value;
    }
  },
  state: new Game(),
  getters: {
    game(state) {
      return state;
    }
  }
};
