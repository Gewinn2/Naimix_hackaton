<template>
  <div class="w-svw flex-grow flex flex-col items-center main-bg">
    <div class="absolute z-10 -top-24 bg-ellipse flex flex-col justify-end items-center">
      <div class="mb-20 text-4xl text-center font-semibold w-full">
        Узнай, что тебя ждёт с помощью <br> Таро-расклада!
      </div>
    </div>
    <div class="w-10/12 h-full z-10 flex flex-col items-center gap-y-2">
      <div class="w-full flex flex-row gap-x-2 items-center mt-4">
        <svg class="w-16 h-16" viewBox="0 0 80 90" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M43.2043 2.12093C41.7424 1.68047 40.2398 2.65641 39.8479 4.30096L26.5088 60.2776C26.1168 61.9228 26.9846 63.6139 28.4471 64.0545L63.2097 74.5284C64.672 74.9687 66.1743 73.9929 66.5663 72.3484L79.9057 16.3715C80.2977 14.7264 79.4297 13.0353 77.9675 12.5947L43.2043 2.12093ZM19.6096 58.1966L28.7318 19.9158L2.03283 27.96C0.570383 28.4006 -0.297505 30.0918 0.0945283 31.7369L13.4337 87.7134C13.8256 89.3578 15.3281 90.3343 16.79 89.8933L51.5529 79.4199C51.5983 79.4064 51.6431 79.3916 51.6873 79.3755L26.5993 71.8168C21.3256 70.2276 18.1958 64.1292 19.6096 58.1966Z" fill="white"/>
        </svg>
        <p class="text-4xl font-semibold">Таро</p>
      </div>
      <div class="flex flex-col self-start text-2xl font-semibold mt-32">
        <p>Для того, чтобы получить расклад, вслух проговори</p>
        <span class="orange-text">“Какое будущее ждёт &lt;Иванова Ивана&gt; в &lt;нашей компании&gt;?” </span>
        <p>Затем выбери 7 карт и получи ответ на вопрос.</p>
      </div>
      <div class="flex flex-col w-full items-center gap-y-10 my-20">
        <div class="flex flex-row relative w-[1396px] h-64">
          <img v-for="_ in 39" :style="{left: (_ * 30) + 'px'}" class="max-w-none min-w-none w-40 h-64 absolute card-animation" src="../assets/cards/back.png"/>
        </div>
        <div class="flex flex-row relative w-[1396px] h-64">
          <img v-for="_ in 39" :style="{left: (_ * 30) + 'px'}" class="max-w-none min-w-none w-40 h-64 absolute card-animation" src="../assets/cards/back.png"/>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">

export default{

  
};

</script>