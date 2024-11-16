<template>
  <div class="flex flex-col gap-y-2 mt-2 max-h-[600px]">
    <p class=" text-lg">{{ title }}</p>
    <div class="flex flex-row mb:self-start rounded-lg input-border">

      <input 
        class=" bg-transparent p-2 outline-none min-w-0 max-w-none w-full flex-grow mb:w-96" 
        type="text" 
        :placeholder
        v-model="searchFilter"
      >

      <div 
        class="w-10 h-10 flex flex-col justify-center items-center cursor-pointer rounded-r-lg bg-slate-200 hover:bg-slate-300 active:bg-slate-400"
        @click="filterUserList"
        >
        <img class="w-7 h-7" src="../assets/icons/icon-search.svg"/>
      </div>

    </div>
    <div class="flex flex-col p-2 pr-0 rounded-lg overflow-hidden input-border">
      <div class="flex-grow flex flex-col gap-y-2 pr-2 scrollable">
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
        <p v-if="filteredUsersList.length === 0">Все сотрудники выбраны!</p>
      </div>
    </div>
    <div class="flex flex-col p-2 pr-0 mt-4 rounded-lg overflow-hidden input-border">
      <div class="flex-grow flex flex-col gap-y-2 pr-2 scrollable">
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
        <p v-if="choosedUsersList.length === 0">Выберите сотрудников</p>
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
