import { defineStore } from "pinia";
import { type StatusCodes, type IStatusWindow } from "@/helpers/constants";

export const useStatusWindowStore = defineStore('statusWindow', {
  state() {
    return{
      statusWindowList: [] as IStatusWindow[],
      statusWindowID: 0,
    }
  },
  actions: {
    showStatusWindow(status: StatusCodes, text: string, time: number): number{
      const newStatusWindowID: number = this.statusWindowID++;
      const newStatusWindow: IStatusWindow = {id: newStatusWindowID, status: status, text: text};
      this.statusWindowList.push(newStatusWindow);

      if(time > 0){
        setTimeout(() => {
          this.deteleStatusWindow(newStatusWindowID);
        }, time);
      }
      return newStatusWindowID;
    },
    deteleStatusWindow(id: number): void{
      const index = this.statusWindowList.findIndex(item => item.id === id);
      if (index !== -1) {
        this.statusWindowList.splice(index, 1);
      }
    }
  }
});