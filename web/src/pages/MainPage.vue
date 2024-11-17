<template>
  <div class="flex flex-row flex-wrap gap-2" ref="imgParent">
    <CardItem v-for="(item, index) in getTarot" :key="item" :id="index" class="w-52"/>
  </div>

</template>
<script lang="ts">

import { TAROT_CARDS } from '@/helpers/constants';
import CardItem from '@/shared/cardItem.vue';

export default{
  components:{
    CardItem,
  },
  data(){
    return{

    }
  },
  computed:{
    getTarot(){
      return TAROT_CARDS;
    },
  },
  methods:{

  }
};

</script>