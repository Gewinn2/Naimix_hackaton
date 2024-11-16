<template>
  <div class="flex flex-grow flex-col items-center justify-center">
    <div class=" py-5 px-7 rounded-lg flex flex-col gap-y-6 login-shadow w-96">
      <div class="w-full text-center text-2xl font-semibold">Авторизация</div>
      <div class="flex flex-col gap-y-5">
        <loginInput 
          type="text" 
          text="Номер телефона или почта" 
          :error="phone.error !== '' || email.error !== ''"
          @input-change="ChangePhoneEmail"
          />
        <loginInput 
          type="password" 
          text="Пароль" 
          :error="password.error !== ''"
          @input-change="ChangePassword"
          />
        <div class=" text-sm text-sky-500 w-full text-right cursor-pointer">Забыли пароль?</div>
      </div>
      <submitButton value="Войти" class="mt-6" @click="SubmitLogin"/>
      <div class="w-full text-center text-sm cursor-default" >
        Ещё не создали аккаунт?
        <span class="text-sky-500 cursor-pointer hover:text-sky-600" @click="$router.push({name: 'SignUpPage'})">Создайте в один клик!</span>
      </div>
    </div>
  </div>
</template>
<script lang="ts">

import loginInput from '@/shared/loginInput.vue';
import submitButton from '@/shared/submitButton.vue';
import { mapStores } from 'pinia';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import { type IValidObject, StatusCodes,type IApiLoginData } from '@/helpers/constants';
import { ValidEmail, ValidPhoneNumber, ValidUserPassword } from '@/helpers/validators';
import { API_Login } from '@/api/api';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      phone: {value: '', error: ''} as IValidObject,
      email: {value: '', error: ''} as IValidObject,
      password: {value: '', error: ''} as IValidObject,
      time: 4000,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore),
  },
  methods:{
    SubmitLogin(){
      if((this.phone.value !== '' || this.email.value !== '') && this.password.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data: IApiLoginData = {
          email: this.email.value,
          phone_number: this.phone.value,
          password: this.password.value,
        };
        API_Login(data)
        .then(response => {
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Авторизация упешна!', this.time);
          this.$router.push({name: 'MainPage'});
        })
        .catch(error => {
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Ошибка авторизации!', this.time);
        });

        return;
      }

      if(this.phone.value === '' && this.email.value === ''){
        if (this.phone.error === '' && this.email.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите номер телефона или электронную почту!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректный номер телефона или адресс электронной почты!', this.time);
      }
      if(this.password.value === '') {
        if(this.password.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите пароль!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.password.error, this.time);
      }
    },
    ChangePhoneEmail(value: string){
       this.phone = ValidPhoneNumber(value); // валидируем
       this.email = ValidEmail(value); // валидируем
       if(value === '') { // если ничего не введено то убираем ошибку
        this.phone.error = '';
        this.email.error = '';
       }
    },
    ChangePassword(value: string){
      this.password = ValidUserPassword(value); // валидируем
      if(value === '') this.password.error = ''; // если ничего не введено то убираем ошибку
    }
  }
}

</script>