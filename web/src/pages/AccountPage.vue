<template>
  <div class="flex flex-col w-full h-full bg-color-ultralight scrollable">
    <div class="flex flex-col justify-evenly items-center gap-y-8 pt-8 md:gap-y-0 md:flex-row md:items-stretch md:pt-16">
      <div class="flex flex-col gap-y-4 md:gap-y-8 items-stretch">
        <p class="text-center text-3xl">Личный Кабинет</p>
        <div class="flex flex-col items-center gap-y-10 mt-8">
          <img class="w-32 h-32" src="../assets/icons/icon-profile.svg"/>
          <div class="flex flex-col gap-y-2 items-center">
            <p class="text-2xl">Иванов Смирнов Соболев</p>
            <p class="text-lg mt-2">Компания: Naimix</p>
            <p class="text-lg">Должность: HR</p>
          </div>
        </div>
        <div class="flex flex-col gap-y-3 items-stretch py-3 px-6 bg-color-light rounded-lg profile-shadow">
          <div class="py-2 px-12 text-center rounded-lg cursor-pointer text-sky-500 bg-white hover:underline header-shadow active:text-sky-300">Сменить пароль</div>
          <div 
            @click="Logout"
            class="py-2 px-12 text-center rounded-lg cursor-pointer text-red-700 bg-white hover:underline header-shadow active:text-red-400">Выйти из аккаунта</div>
        </div>
        <!--temporary-->
        <div class=" flex flex-row gap-x-2">
          <input 
            type="text" 
            v-model="companyName" 
            class="py-2 px-4 rounded-lg bg-slate-50 text-gray-800 text-lg font-medium max-w-none w-[400px]" 
            placeholder="Введите название компании"
            />
          <div @click="createCompany" class="py-2 px-4 text-xl rounded-lg font-medium btn">
            Создать компанию
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { StatusCodes } from '@/helpers/constants';
import { API_LogOut, API_CreateCompany } from '@/api/api';
export default{
  data(){
    return{
      time: 3000,
      companyName: '',
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
    },
    
    createCompany(){
      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Создаем компанию...', -1);

      API_CreateCompany(this.companyName)
      .then(response => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Компания успешно создана!', 3000);
      })
      .catch(error => {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так пр создании компании!', 3000);
      })
      .finally(() => {
        this.statusWindowStore.deteleStatusWindow(stID);
      })
    }
  }
};

</script>