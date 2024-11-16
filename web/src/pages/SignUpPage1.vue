<template>
  <div class="flex flex-grow flex-col items-center justify-center">
    <div class="py-5 px-7 rounded-lg flex flex-col gap-y-6 login-shadow w-96">
      <div class="w-full text-center text-2xl font-semibold">Регистрация</div>
      <div class="flex flex-col gap-y-5">
        <loginInput 
          type="email" 
          text="Почта"
          :start-text="loginStore.SignupEmail"
          :error="email.error !== ''" 
          @input-change="ChangeEmail"
          />
        <loginInput 
          type="tel" 
          text="Телефон" 
          :start-text="loginStore.SignupPhoneNumber"
          :error="phone.error !== ''" 
          @input-change="ChangePhone"
          />
        <loginInput 
          type="password" 
          text="Пароль" 
          :start-text="loginStore.SignupPassword"
          :error="password.error !== ''" 
          @input-change="ChangePassword"
          />
      </div>
      <submitButton value="Далее" class="mt-8" @click="nextPage"/>
      <div class="w-full text-center text-sm cursor-default">
        Уже есть аккаунт?
        <span class="text-sky-500 cursor-pointer hover:text-sky-600" @click="$router.push({name: 'LoginPage'})">Войдите в аккаунт!</span>
      </div>
    </div>
  </div>
</template>
<script lang="ts">

import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { useLoginStore } from '@/stores/loginStore';
import loginInput from '@/shared/loginInput.vue';
import submitButton from '@/shared/submitButton.vue';
import { StatusCodes, type IValidObject } from '@/helpers/constants';
import { ValidEmail, ValidPhoneNumber, ValidUserPassword } from '@/helpers/validators';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      email: {value: '', error: ''} as IValidObject,
      phone: {value: '', error: ''} as IValidObject,
      password: {value: '', error: ''} as IValidObject,
      time: 4000,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useLoginStore),
  },
  methods:{
    nextPage(){
      if (this.email.value !== '' && this.phone.value !== '' && this.password.value !== ''){
        this.$router.push({name: 'SignUpPage2'});
        return;
      }
      if(this.email.value === '') {
        if(this.email.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите адрес электронной почты!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.email.error, this.time);
      }
      if(this.phone.value === '') {
        if(this.phone.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите номер телефона!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.phone.error, this.time);
      }
      if(this.password.value === '') {
        if(this.password.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error, this.time);
      }
    },
    ChangeEmail(value: string){
      this.email = ValidEmail(value); // валидируем
      if(this.email.error === '') this.loginStore.SignupEmail = this.email.value; // если корректно, то запоминаем
      if(value === '') this.email.error = ''; // если ничего не введено то убираем ошибку
    },
    ChangePhone(value: string){
      this.phone = ValidPhoneNumber(value); // валидируем
      if(this.phone.error === '') this.loginStore.SignupPhoneNumber = this.phone.value; // если корректно, то запоминаем
      if(value === '') this.phone.error = ''; // если ничего не введено то убираем ошибку
    },
    ChangePassword(value: string){
      this.password = ValidUserPassword(value); // валидируем
      if(this.password.error === '') this.loginStore.SignupPassword = this.password.value; // если корректно, то запоминаем
      if(value === '') this.password.error = ''; // если ничего не введено то убираем ошибку
    }
  }
}
</script>