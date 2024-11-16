import { defineStore } from "pinia";
import { type IPickCard, type ICard } from "@/helpers/constants";

export const useTarotStore = defineStore('tarot', {
  state(){
    return{
      cards: [] as IPickCard[],

      categoryCards: [] as ICard[],
    }
  },
  actions:{
    generateCards(): void{
      for(let i = 0; i < 78; i++) this.cards.push({isPicked: false});
    },
  }
});