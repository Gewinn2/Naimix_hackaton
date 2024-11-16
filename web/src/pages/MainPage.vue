<template>
  <div class="flex flex-row flex-wrap gap-2">
    <img v-for="(item, index) in getTarot" :key="item" :src="'src'+getCard(index)" class="w-52 aspect-auto border-solid border border-gray-800"/>
  </div>

</template>
<script lang="ts">
import { GET_CARD_IMAGE, TAROT_CARDS } from '@/helpers/constants';

export default{
  data(){
    return{

    }
  },
  computed:{
    getTarot(){
      return TAROT_CARDS;
    }
  },
  methods:{
    getCard(cardID: number):string {
      return GET_CARD_IMAGE(cardID);
    }
  }
};

</script>