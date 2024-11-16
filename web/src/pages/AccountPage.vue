<template>
  <div class="flex flex-col w-full h-full bg-color-ultralight scrollable">
    <div class="flex flex-col justify-evenly items-center gap-y-8 pt-8 md:gap-y-0 md:flex-row md:items-stretch md:pt-16">
      <div class="flex flex-col gap-y-4 md:gap-y-8 items-stretch">
        <p class="text-center text-3xl">Личный Кабинет</p>
        <div class="flex flex-col items-center gap-y-10 mt-8">
          <img class="w-32 h-32" src="../assets/icons/icon-profile.svg"/>
          <div class="flex flex-col gap-y-2 items-center">
            <p class="text-2xl">Иванов Смирнов Соболев</p>
            <p class="text-lg mt-2">Компания: ООО Псковский туберкулез</p>
            <p class="text-lg">Должность: слесарь-автомеханик</p>
          </div>
        </div>
        <div class="flex flex-col gap-y-3 items-stretch py-3 px-6 bg-color-light rounded-lg profile-shadow">
          <div class="py-2 px-12 text-center rounded-lg cursor-pointer text-sky-500 bg-white hover:underline header-shadow active:text-sky-300">Сменить пароль</div>
          <div 
            @click="Logout"
            class="py-2 px-12 text-center rounded-lg cursor-pointer text-red-700 bg-white hover:underline header-shadow active:text-red-400">Выйти из аккаунта</div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { StatusCodes } from '@/helpers/constants';
import { API_LogOut } from '@/api/api';
export default{
  data(){
    return{
      time: 3000,
    }
  },
  computed:{
    ...mapStores(useStatusWindowStore),
  },
  methods:{
    Logout(){
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
      API_LogOut()
      .then(response => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Вы вышли из аккаунта!', this.time);
        this.$router.push({name: 'LoginPage'});
      })
      .catch(error => {
        this.statusWindowStore.deteleStatusWindow(stID);
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при выходе из аккаунта!', this.time);
      })
    }
  }
};

</script>