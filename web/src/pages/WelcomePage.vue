<template>
  <div class="flex flex-col flex-grow w-full items-center welcome-bg-image">
    <div class="flex flex-row justify-between items-stretch w-11/12 h-full">
      <div class="flex flex-col items-start w-[700px] gap-y-6">
        <p class="text-5xl font-medium my-10 mb-2">Привет!</p>
        <p class="text-5xl font-medium">Ты - HR крупной компании и столкнулся с трудностями при отборе кандидатов?</p>
        <p class="text-5xl font-medium">Или, может, ты только устраиваешься на работу и хочешь заглянуть в будущее?</p>
        <p class="text-5xl font-medium">Узнай, какие планы на ваше сотрудничество у судьбы!</p>
      </div>
      <div class="flex flex-col justify-end w-[800px] gap-y-6">
        <p class="text-5xl font-medium text-right">Войди в свой аккаунт</p>
        <p class="text-5xl font-medium text-right">или зарегистрируй новый,</p>
        <p class="text-5xl font-medium text-right">если ты впервые</p>
        <p class="text-5xl font-medium text-right">на нашем сайте</p>
        <div @click="$router.push({name: 'LoginPage'})" class="py-3 px-6 text-2xl font-bold self-end rounded-2xl my-10 yellow-text cursor-pointer bg-slate-50 hover:bg-slate-100 active:bg-slate-200">
          Войти
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
export default{

};
</script>