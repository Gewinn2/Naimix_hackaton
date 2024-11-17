<template>
  <div class="flex flex-row gap-x-5 min-h-20">
    <CardItem :id="cardId" class="w-52" :class="{'rotate-180' : reversed}"/>
    <div class="flex flex-col items-start gap-y-1">
      <p class="text-3xl">{{ scope }}</p>
      <p class="text-2xl orange-text">{{ category }}</p>
      <p class="text-xl">{{ description }}</p>
      <p class="text-lg text-justify">{{ text }}</p>
    </div>
  </div>
</template>
<script lang="ts">
import CardItem from './cardItem.vue';

export default{
  components:{
    CardItem,
  },
  props:{
    cardId: {
      type: Number,
      required: true,
    },
    scope: {
      type: String,
      required: true,
    },
    reversed: {
      type: Boolean,
      required: false,
      default: false,
    },
    category: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
  }
};
</script>