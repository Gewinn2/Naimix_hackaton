<template>
  <StatusWindow />
  <Header />
  <RouterView />
</template>
<style v-global>
.list-enter-active,
.list-leave-active {
  transition: all 0.5s ease;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>
<script lang="ts">

import { mapStores } from "pinia";
import { useStatusWindowStore } from "./stores/statusWindowStore";
import StatusWindow from "./entities/statusWindow.vue";
import Header from "./entities/header.vue";

export default {
  components: {
    StatusWindow,
    Header,
  },
  data() {
    return {

    };
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods: {

  },
};
</script>
