<template>
  <div class="grid grid-cols-12 py-2 px-4 rounded-lg input-border">
    <p class="col-span-8 text-lg">{{ name }}</p>
    <p class=" col-span-3 text-lg">{{ role }}</p>
    <input class="col-span-1" type="checkbox" v-model="check">
  </div>
</template>
<script lang="ts">
export default{
  emits: ['changeChoose'],
  props: {
    id:{
      type: Number,
      required: true,
    },
    name: {
      type: String,
      required: true,
    },
    role: {
      type: String,
      required: false,
      default: '',
    },
    isChoosed: {
      type: Boolean,
      required: false,
      default: false,
    },
  },
  data(){
    return{
      check: this.isChoosed,
    }
  },
  watch:{ 
    check:{
      handler(val){
        this.$emit('changeChoose', this.id, this.check);
      },
      immediate: false,
    }
  },
};
</script>