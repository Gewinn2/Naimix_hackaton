<template>
  <div class="flex flex-col items-start relative">
    <label
      class="absolute transition-all text-gray-800"
      :class="{ 'login-label': !focusLabel, 'login-label-focus': focusLabel }"
      @click="$refs.input.focus()"
      >{{ text }}</label>

    <div
      class="flex flex-row items-end gap-x-2 w-full h-12 border-bottom  transition-input hover:border-sky-500"
      :class="{ 'border-sky-500': isFocused, 'border-slate-500': !isFocused, 'bad-input': error}"
    >
      <input
        class="bg-transparent outline-none h-7 pl-1 text-lg flex-1 text-gray-800"
        :type="inputType"
        v-model="inputValue"
        ref="input"
        @focusin="onFocus"
        @focusout="unFocus"
      />

      <div
        v-if="type === 'password'"
        @mousedown="onHold"
        @mouseup="unHold"
        @touchstart="onHold"
        @touchend="unHold"
        @dragend="unHold"
      >
        <img
          v-if="inputType !== type"
          src="../assets/icons/icon-password-hide.svg"
          alt="Показать пароль"
          class="w-6 h-6 cursor-pointer mr-1"
        />
        <img
          v-else
          src="../assets/icons/icon-password-show.svg"
          class="w-6 h-6 cursor-pointer mr-1"
        />
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default {
  props: {
    type: {
      type: String,
      required: true,
    },
    text: {
      type: String,
      required: true,
    },
    startText:{
      type: String,
      required: false,
      default: '',
    },
    error: {
      type: Boolean,
      required: false,
      default: false,
    }
  },
  emits: ['inputChange'],
  data() {
    return {
      inputType: this.type,
      isFocused: false,
      inputValue: this.startText,
    };
  },
  computed: {
    focusLabel() {
      return this.isFocused || this.inputValue !== "" || this.type === 'date';
    },
  },
  methods: {
    onHold() {
      this.inputType = "text";
    },
    unHold() {
      this.inputType = this.type;
    },
    onFocus() {
      this.isFocused = true;
    },
    unFocus() {
      this.isFocused = false;
    },
  },
  watch:{
    'inputValue':{
      handler(val){
        this.$emit('inputChange', val);
      },
      immediate: true,
    }
  }
};
</script>
