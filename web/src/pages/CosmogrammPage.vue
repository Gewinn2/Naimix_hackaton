<template>
  <div class="w-full h-full flex flex-col items-center overflow-y-scroll">
    <div class="w-8/12 h-full flex flex-col z-20 relative top-5">
      <div class="absolute z-10 right-0 flex flex-col justify-center items-center">
        <svg class="absolute w-[500px] h-[300px] z-10" viewBox="0 0 766 404" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="250" cy="154" r="250" fill="#9A9A9A"/>
          <circle cx="590" cy="154" r="250" fill="#FF9929"/>
        </svg>
        <p class="text-2xl font-medium text-center z-20 relative">Воспользуйся нашей космограммой,<br> чтобы провести анализ совместимости<br> рекрута и будущих коллег.</p>
      </div>
      <section>
        <div class="flex flex-row items-center self-start w-full mt-5 mb-10 gap-x-2">
          <svg class="w-16 h-16" viewBox="0 0 60 70" fill="none" xmlns="http://www.w3.org/2000/svg">
            <g clip-path="url(#clip0_77_379)">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M4.8777 1.90252C6.3441 -0.265109 9.03904 -0.635434 10.897 1.07537C15.2899 5.12036 18.8019 10.3185 21.1471 16.2466C22.8684 20.5978 23.9254 25.252 24.2823 30.0002H35.7182C36.075 25.252 37.132 20.5978 38.8533 16.2466C41.1985 10.3185 44.7107 5.12036 49.1036 1.07537C50.9614 -0.635434 53.6563 -0.265109 55.1229 1.90252C56.589 4.07014 56.2719 7.21426 54.4136 8.92506C51.0806 11.9944 48.4153 15.9388 46.6359 20.4371C45.4354 23.4717 44.6606 26.7004 44.3331 30.0002H55.7143C58.0813 30.0002 60 32.2388 60 35.0002C60 37.7616 58.0813 40.0002 55.7143 40.0002H44.3331C44.6606 43.3001 45.4354 46.5289 46.6359 49.5635C48.4153 54.0617 51.0806 58.0062 54.4136 61.0757C56.2719 62.7862 56.589 65.9302 55.1229 68.0982C53.6563 70.2657 50.9614 70.6362 49.1036 68.9252C44.7107 64.8802 41.1985 59.6822 38.8533 53.7542C37.132 49.4028 36.075 44.7485 35.7182 40.0002H24.2823C23.9255 44.7485 22.8685 49.4028 21.1471 53.7542C18.8019 59.6822 15.2899 64.8802 10.897 68.9252C9.03904 70.6362 6.3441 70.2657 4.8777 68.0982C3.41129 65.9302 3.72871 62.7862 5.58669 61.0757C8.92007 58.0062 11.585 54.0617 13.3646 49.5635C14.5651 46.5289 15.3397 43.3001 15.6675 40.0002H4.28571C1.91878 40.0002 0 37.7616 0 35.0002C0 32.2388 1.91878 30.0002 4.28571 30.0002H15.6674C15.3397 26.7004 14.565 23.4717 13.3646 20.4371C11.585 15.9388 8.92007 11.9944 5.58669 8.92506C3.72871 7.21426 3.41129 4.07014 4.8777 1.90252Z" fill="white"/>
            </g>
            <defs>
            <clipPath id="clip0_77_379">
            <rect width="60" height="70" fill="white"/>
            </clipPath>
            </defs>
          </svg>
          <p class="text-4xl font-semibold">Космограмма</p>
        </div>
        <UserList 
          class="w-full mb-2"
          title="Выберите от 1 до 10 коллег для анализа совместимости" 
          placeholder="Введите ФИО" 
          :item-component="getUserItem"/>
      </section>
      <section>
        <div class="flex flex-row gap-x-2 text-lg items-center py-2">
          <p>Введите год рождения и ФИО рекрута: </p>
          <div>
            <input 
              type="date" 
              v-model="recrutBirthDay"
              class="outline-none py-2 px-4 bg-slate-50 text-lg text-gray-800 font-medium border-2 rounded-lg border-solid border-gray-400"/>
          </div>
          <div>
            <input 
              type="text" 
              v-model="recrutFIO"
              class="outline-none py-2 px-4 bg-slate-50 text-lg text-gray-800 font-medium border-2 rounded-lg border-solid border-gray-400 max-w-none w-[350px]" placeholder="ФИО рекрута"/>
          </div>
          <div v-if="1 <= choosedUsers.length && choosedUsers.length <= 10" @click="analyseCosmogramm" class="flex flex-row items-center justify-center py-2 px-4 rounded-lg cursor-pointer btn text-slate-50 text-gl font-medium">
            Анализ совместимости
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useCosmogrammStore } from '@/stores/cosmogrammStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import UserList from '@/entities/userList.vue';
import { type IUsersList, type ICosmogrammRequest, minDate, maxDate, StatusCodes } from '@/helpers/constants';
import UserItem from '@/shared/userItem.vue';
import { API_Cosmogramm } from '@/api/api';
export default{
  components:{
    UserList,
    UserItem,
  },
  data(){
    return{
      recrutBirthDay: '',
      recrutFIO: '',
    }
  },
  computed:{
    ...mapStores(useCosmogrammStore, useStatusWindowStore),

    getUserItem(){
      return UserItem;
    },

    choosedUsers(){
      return this.cosmogrammStore.usersList.filter(item => item.isChoosed);
    }
  },
  methods: {
    analyseCosmogramm(){
      const choosedUser = this.cosmogrammStore.usersList.filter(item => item.isChoosed);
      if(choosedUser.length < 1 || choosedUser.length > 10) return;
      const recrutBD = new Date(this.recrutBirthDay);
      if(recrutBD < minDate){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Система не поддерживает пользователей, родившихся до 1950г!', 3000);
        return;
      }
      if(recrutBD > maxDate){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Система не поддерживает пользователей, родившихся после 2011г!', 3000);
        return;
      }
      const recrutFIO = this.recrutFIO.split(' '); 
      if(recrutFIO.length < 2 || recrutFIO.length > 3){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректное ФИО рекрута!', 3000);
        return;
      }

      const data: ICosmogrammRequest = {  
        birth_date: this.recrutBirthDay,
        name: recrutFIO[0],
        surname: recrutFIO[1],
        third_name: recrutFIO[2] ? recrutFIO[2] : '',
      };

      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Создаем космограмму совместимости...', -1);
      API_Cosmogramm(data)
      .then(response => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Космограмма готова!', 3000);
        //
      })
      .catch(error => {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при создании космограммы!', 3000);
        //
      })
      .finally(() => {
        this.statusWindowStore.deteleStatusWindow(stID);
      });
    }
  }
};
</script>