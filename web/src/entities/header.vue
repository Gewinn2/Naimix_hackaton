<template>
  <div class="flex flex-row flex-shrink-0 items-stretch h-14 w-full px-5 z-40 header-bg">
    <div class="h-full grid content-center mx-16">
      <div class="font-semibold text-3xl cursor-pointer select-none" @click="$router.push({name: 'MainPage'})">Harmony Compass</div>
    </div>
    <div class="h-full flex flex-row gap-x-4 flex-grow text-lg font-semibold">
      <div 
        @click="$router.push({name: 'TarotPage'})"
        :class="{'border-b-sky-500' : headerStore.currentPage === 'TarotPage'}"
        class="h-full min-w-16 grid cursor-pointer justify-center content-center select-none border-bottom-header border-gray-500 hover:border-b-sky-500">
        <a>ТАРО</a>
      </div>
      <div 
        @click="$router.push({name: 'CosmogrammPage'})"
        :class="{'border-b-sky-500' : headerStore.currentPage === 'CosmogrammPage'}"
        class="h-full min-w-16 grid cursor-pointer justify-center content-center select-none border-bottom-header border-gray-500 hover:border-b-sky-500">
        <a>КОСМОГРАММА</a>
      </div>
    </div>
    <div class="flex flex-row gap-x-2 items-center mr-16 cursor-pointer" @click="$router.push({name: 'AccountPage'})">
      <img class="w-10 h-10" src="../assets/icons/icon-profile.svg"/>
      <div class="text-lg">UserName</div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useHeaderStore } from '@/stores/headerStore';
export default {
  computed:{
    ...mapStores(useHeaderStore),
  },
};
</script>
