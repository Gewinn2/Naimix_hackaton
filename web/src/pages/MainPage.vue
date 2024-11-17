<template>
  <div class="flex flex-col w-full items-center main-bg-image mt-2 flex-grow">
    <div class="flex flex-col w-full xl:w-10/12 items-center gap-y-20">
      <p class="text-5xl font-bold w-[1000px] text-wrap mt-24 text-center">С помощью Harmony Compass ты сможешь проанализровать дальнейшие перспективы сотрудничества с помощью:</p>
      <div class="flex flex-row justify-evenly w-full">
        
        <div class="flex flex-col items-center relative justify-center w-[450px] h-[300px] cursor-pointer" @click="$router.push({name: 'TarotPage'})">
          <svg class="absolute w-[450px] h-[300px] z-10" viewBox="0 0 450 300" fill="none" xmlns="http://www.w3.org/2000/svg">
            <ellipse cx="225" cy="150" rx="225" ry="150" fill="white"/>
          </svg>
          <div class="flex flex-col justify-center items-center gap-y-2 z-20">
            <svg class="w-24 h-24" viewBox="0 0 80 88" fill="none" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M43.2042 0.120986C41.7422 -0.319474 40.2397 0.656467 39.8478 2.30102L26.5087 58.2777C26.1166 59.9228 26.9845 61.6139 28.447 62.0545L63.2096 72.5284C64.6719 72.9688 66.1742 71.9929 66.5662 70.3485L79.9056 14.3716C80.2976 12.7264 79.4296 11.0353 77.9673 10.5947L43.2042 0.120986ZM19.6094 56.1967L28.7317 17.9159L2.03271 25.96C0.570261 26.4006 -0.297627 28.0918 0.0944063 29.7369L13.4336 85.7134C13.8254 87.3578 15.328 88.3343 16.7899 87.8933L51.5528 77.4199C51.5981 77.4064 51.643 77.3916 51.6872 77.3755L26.5992 69.8168C21.3254 68.2277 18.1957 62.1292 19.6094 56.1967Z" fill="#353535"/>
            </svg>
            <p class="text-5xl text-gray-800 font-bold">Таро</p>
          </div>
        </div>

        <div class="flex flex-col items-center relative justify-center w-[450px] h-[300px] cursor-pointer" @click="$router.push({name: 'CosmogrammPage'})">
          <svg class="absolute w-[450px] h-[300px] z-10" viewBox="0 0 450 300" fill="none" xmlns="http://www.w3.org/2000/svg">
            <ellipse cx="225" cy="150" rx="225" ry="150" fill="white"/>
          </svg>
          <div class="flex flex-col justify-center items-center gap-y-2 z-20">
            <svg class="w-24 h-24" viewBox="0 0 60 70" fill="none" xmlns="http://www.w3.org/2000/svg">
              <g clip-path="url(#clip0_41_175)">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M4.8777 1.90252C6.3441 -0.265109 9.03904 -0.635434 10.897 1.07537C15.2899 5.12036 18.8019 10.3185 21.1471 16.2466C22.8684 20.5978 23.9254 25.252 24.2823 30.0002H35.7182C36.075 25.252 37.132 20.5978 38.8533 16.2466C41.1985 10.3185 44.7107 5.12036 49.1036 1.07537C50.9614 -0.635434 53.6563 -0.265109 55.1229 1.90252C56.589 4.07014 56.2719 7.21426 54.4136 8.92506C51.0806 11.9944 48.4153 15.9388 46.6359 20.4371C45.4354 23.4717 44.6606 26.7004 44.3331 30.0002H55.7143C58.0813 30.0002 60 32.2388 60 35.0002C60 37.7616 58.0813 40.0002 55.7143 40.0002H44.3331C44.6606 43.3001 45.4354 46.5289 46.6359 49.5635C48.4153 54.0617 51.0806 58.0062 54.4136 61.0757C56.2719 62.7862 56.589 65.9302 55.1229 68.0982C53.6563 70.2657 50.9614 70.6362 49.1036 68.9252C44.7107 64.8802 41.1985 59.6822 38.8533 53.7542C37.132 49.4028 36.075 44.7485 35.7182 40.0002H24.2823C23.9255 44.7485 22.8685 49.4028 21.1471 53.7542C18.8019 59.6822 15.2899 64.8802 10.897 68.9252C9.03904 70.6362 6.3441 70.2657 4.8777 68.0982C3.41129 65.9302 3.72871 62.7862 5.58669 61.0757C8.92007 58.0062 11.585 54.0617 13.3646 49.5635C14.5651 46.5289 15.3397 43.3001 15.6675 40.0002H4.28571C1.91878 40.0002 0 37.7616 0 35.0002C0 32.2388 1.91878 30.0002 4.28571 30.0002H15.6674C15.3397 26.7004 14.565 23.4717 13.3646 20.4371C11.585 15.9388 8.92007 11.9944 5.58669 8.92506C3.72871 7.21426 3.41129 4.07014 4.8777 1.90252Z" fill="#353535"/>
              </g>
              <defs>
              <clipPath id="clip0_41_175">
              <rect width="60" height="70" fill="white"/>
              </clipPath>
              </defs>
            </svg>
            <p class="text-5xl text-gray-800 font-bold">Космограмма</p>
          </div>
        </div>

      </div>
    </div>
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