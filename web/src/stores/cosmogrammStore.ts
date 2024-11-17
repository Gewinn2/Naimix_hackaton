import { defineStore } from "pinia";
import { type IUsersList, type ICosmogrammPlanet, type ICosmogrammCompatibilitiy } from "@/helpers/constants";

export const useCosmogrammStore = defineStore('cosmogramm', {
  state(){
    return{
      usersList: [
        {id: 1, name: 'Иванов Иван Иванович', role: 'Слесарь', isChoosed: false},
        {id: 2, name: 'Николаев Николай Николаевич', role: 'Слесарь', isChoosed: false},
        {id: 3, name: 'Зубенко Михаил Петрович', role: 'Сисадмин', isChoosed: false},
      ] as IUsersList[],

      cosmoPlanets: [] as ICosmogrammPlanet[],

      compatibility: [] as ICosmogrammCompatibilitiy[],
    }
  },
});