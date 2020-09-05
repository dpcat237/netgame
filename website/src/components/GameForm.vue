<template>
  <div>
    <div v-if="!isSelected">
      <v-btn @click="observe">
        Observe
      </v-btn>
      <v-btn class="ml-10" @click="isSelected = true">
        Play
      </v-btn>
    </div>
    <div v-else>
      <v-form class="request-form">
        <v-text-field v-model="name" :clearable="true" label="Name" required />
        <v-spacer class="mb-5" />
        <v-slider
          class="number"
          style="float: left;"
          v-model="numberOne"
          thumb-label="always"
          :max="10"
          :min="1"
        ></v-slider>
        <v-slider
          class="number"
          v-model="numberTwo"
          thumb-label="always"
          :max="10"
          :min="1"
        />
        <v-btn class="ml-10" @click="submitRequest">
          Submit
        </v-btn>
      </v-form>
      <h4 class="mt-5 error-message">{{ singError }}</h4>
    </div>
  </div>
</template>

<script>
export default {
  name: "Playground",
  data() {
    return {
      isSelected: false,
      name: "",
      numberOne: 1,
      numberTwo: 2,
      singError: ""
    };
  },
  computed: {
    errorMessage: function() {
      return this.$store.getters.game.errorMessage;
    },
    playerName: function() {
      return this.$store.getters.game.name;
    },
    requestFinished: function() {
      return this.$store.getters.game.signedRequestFinished;
    }
  },
  watch: {
    requestFinished: function(newValue) {
      if (!newValue) {
        return;
      }

      if (this.errorMessage === "" && this.playerName !== "") {
        this.$emit("selected");
        return;
      }

      if (this.errorMessage !== "") {
        this.singError = this.errorMessage;
        return;
      }
      this.singError = "Error signing. Try again a few minutes.";
    }
  },
  methods: {
    submitRequest() {
      this.singError = "";
      this.$store.dispatch("sign", {
        name: this.name,
        numberOne: this.numberOne,
        numberTwo: this.numberTwo
      });
      this.$store.dispatch("observe");
    },
    observe() {
      this.$store.dispatch("observe");
      this.$emit("selected");
    }
  }
};
</script>

<style scoped lang="scss">
.request-form {
  width: 500px;
  margin: auto;

  .number {
    width: 240px;
  }
}

.error-message {
  color: red;
}
</style>
