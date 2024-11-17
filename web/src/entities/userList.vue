<template>
  <div class="flex flex-col gap-y-4 mt-2 max-h-[600px]">
    <p class=" text-3xl font-semibold w-[400px]">{{ title }}</p>
    <div class="flex flex-row mb:self-start rounded-lg input-border overflow-hidden">

      <input 
        class=" bg-slate-50 p-3 outline-none min-w-0 max-w-none w-full flex-grow mb:w-96 text-2xl font-medium text-gray-900" 
        type="text" 
        :placeholder
        v-model="searchFilter"
      >

      <div 
        class="w-14 h-14 flex flex-col justify-center items-center cursor-pointer bg-slate-200 hover:bg-slate-300 active:bg-slate-400"
        @click="filterUserList"
        >
        <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
      </div>

    </div>
    <div class="flex flex-col p-3 rounded-lg overflow-hidden input-border">
      <div class="flex-grow flex flex-col gap-y-3">
        <component 
          v-for="item in filteredUsersList" 
          :key = item.id
          :is = itemComponent
          :id = item.id 
          :name = item.name
          :role = item.role 
          :isChoosed = item.isChoosed
          @changeChoose = "choosedUser"
        /> 
        <p v-if="filteredUsersList.length === 0" class="text-lg">Все сотрудники выбраны!</p>
      </div>
    </div>
    <div class="flex flex-col p-3 mt-4 rounded-lg overflow-hidden input-border">
      <div class="flex-grow flex flex-col gap-y-3">
        <component 
          v-for="item in choosedUsersList" 
          :key = item.id
          :is = itemComponent
          :id = item.id 
          :name = item.name
          :role = item.role 
          :isChoosed = item.isChoosed
          @changeChoose = "choosedUser"
        /> 
        <p v-if="choosedUsersList.length === 0" class="text-lg">Выберите сотрудников</p>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useCosmogrammStore } from '@/stores/cosmogrammStore';
import { type IUsersList } from '@/helpers/constants';

export default{
  props: {
    title: {
      type: String,
      required: true,
    },
    placeholder:{
      type: String,
      required: false,
      default: '',
    },
    itemComponent: {
      required: true,
    },
  },
  data(){
    return{
      searchFilter: '',
      filteredUsersList: [] as IUsersList[],
      choosedUsersList: [] as IUsersList[],
    }
  },
  computed:{
    ...mapStores(useCosmogrammStore),
  },
  mounted(){
    this.filterUserList();
  },
  methods:{
    filterUsersListInput(){
      if(this.searchFilter === '' || this.cosmogrammStore.usersList.length > 1000) return this.filteredUsersList;
      this.filterUserList();
    },
    filterUserList(){
      this.filteredUsersList = [];
      this.choosedUsersList = [];
      for(let item of this.cosmogrammStore.usersList){
        if(item.isChoosed) this.choosedUsersList.push(item);
        else{
          if(this.searchFilter === '') this.filteredUsersList.push(item);

          else if(item.name.toLowerCase().includes(this.searchFilter.toLowerCase())){
            this.filteredUsersList.push(item);
          }
        }
      }
    },
    choosedUser(id: number, value: boolean){
      for(let item of this.cosmogrammStore.usersList) {
        if(item.id === id){
          item.isChoosed = value;
          break;
        }
      }
      this.filterUserList();
    }
  },
  watch:{
    searchFilter(val){
      if(val === '') {
        this.filteredUsersList = this.cosmogrammStore.usersList.filter(item => !item.isChoosed);
        return;
      }
      if(this.cosmogrammStore.usersList.length > 1000) return;
      else this.filterUserList();
    }
  }
};

</script>
