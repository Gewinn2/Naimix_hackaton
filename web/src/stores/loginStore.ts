import { defineStore } from "pinia";

export const useLoginStore = defineStore('login', {
  state(){
    return{
      SignupName: '',
      SignupSurname: '',
      SignupThirdname: '',
      SignupEmail: '',
      SignupPhoneNumber: '',
      SignupPassword: '',
      SignupBirthDate: '',
    }
  },
});