<template>
  <div class="absolute self-center top-16 flex flex-col gap-y-1 z-50">
    <TransitionGroup name="list">
      <div
        v-for="item in statusWindowStore.statusWindowList"
        :key="item.id"
        class="flex flex-row gap-x-2 items-center py-1 px-3 rounded-md"
        :class="{
          'bg-amber-500': isLoading(item.status),
          'bg-green-600': isSuccess(item.status),
          'bg-red-500': isError(item.status),
          'bg-cyan-600': isInfo(item.status),
        }"
      >
        <div>
          <img
            v-if="isError(item.status)"
            class="w-7 h-7"
            src="../assets/icons/icon-error.svg"
          />
          <img
            v-else-if="isSuccess(item.status)"
            class="w-7 h-7"
            src="../assets/icons/icon-success.svg"
          />
          <img
            v-else-if="isLoading(item.status)"
            class="w-7 h-7"
            src="../assets/icons/icon-loading.svg"
          />
          <img
            v-else-if="isInfo(item.status)"
            class="w-7 h-7"
            src="../assets/icons/icon-info.svg"
          />
        </div>
        <div class="text-slate-50">
          {{ item.text }}
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>
<script lang="ts">

import { mapStores } from "pinia";
import { useStatusWindowStore } from "../stores/statusWindowStore";
import { StatusCodes } from "@/helpers/constants";

export default {
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods: {
    isLoading(status: StatusCodes): Boolean{
      return status === StatusCodes.loading;
    },
    isSuccess(status: StatusCodes): Boolean{
      return status === StatusCodes.success;
    },
    isError(status: StatusCodes): Boolean{
      return status === StatusCodes.error;
    },
    isInfo(status: StatusCodes): Boolean{
      return status === StatusCodes.info;
    },
  }
};
</script>
