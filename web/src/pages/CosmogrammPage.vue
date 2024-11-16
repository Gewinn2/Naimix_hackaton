<template>
  <div class="w-full h-full flex flex-col items-center overflow-y-scroll">
    <div class="w-10/12 h-full flex flex-col">
      <section>
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
            <input type="date" class=" outline-none py-2 px-4 bg-transparent border-2 rounded-lg border-solid border-gray-400"/>
          </div>
          <div>
            <input type="text" class="  outline-none py-2 px-4 bg-transparent border-2 rounded-lg border-solid border-gray-400" placeholder="ФИО рекрута"/>
          </div>
          <div v-if="1 <= choosedUsers.length && choosedUsers.length <= 10" class="flex flex-row items-center justify-center py-2 px-4 rounded-lg cursor-pointer btn text-slate-50">
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
import UserList from '@/entities/userList.vue';
import { type IUsersList } from '@/helpers/constants';
import UserItem from '@/shared/userItem.vue';
export default{
  components:{
    UserList,
    UserItem,
  },
  computed:{
    ...mapStores(useCosmogrammStore),

    getUserItem(){
      return UserItem;
    },

    choosedUsers(){
      return this.cosmogrammStore.usersList.filter(item => item.isChoosed);
    }
  }
};
</script>