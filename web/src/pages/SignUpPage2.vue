<template>
  <div class="flex flex-grow flex-col items-center justify-center">
    <div class="py-5 px-7 rounded-lg flex flex-col gap-y-6 login-shadow w-96">
      <div class="p-2 rounded-lg absolute cursor-pointer bg-slate-200" @click="$router.push({name: 'SignUpPage1'})">
        <img class=" w-6 h-6" src="../assets/icons/icon-back.svg"/>
      </div>
      <div class="w-full text-center text-2xl font-semibold">Регистрация</div>
      <div class="flex flex-col gap-y-5">
        <loginInput 
          type="text" 
          text="Фамилия" 
          :start-text="loginStore.SignupSurname"
          :error="surname.error !== ''" 
          @input-change="ChangeSurname"
          />
        <loginInput 
          type="text" 
          text="Имя" 
          :start-text="loginStore.SignupName"
          :error="name.error !== ''" 
          @input-change="ChangeName"
          />
        <loginInput 
          type="text" 
          text="Отчество" 
          :start-text="loginStore.SignupThirdname"
          :error="thirdname.error !== ''" 
          @input-change="ChangeThirdname"
          />
        <loginInput 
          type="date" 
          text="Дата рождения" 
          :start-text="loginStore.SignupBirthDate"
          :error="birthDate.error !== ''" 
          @input-change="ChangeBirthDate"
          />
      </div>
      <submitButton value="Зарегистрироваться" class="mt-8" @click="SubmitSignUp"/>
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
import { StatusCodes, type IValidObject, type IApiSignUpData } from '@/helpers/constants';
import { ValidUserName, ValidUserSurname, ValidUserThirdname, ValidBirthDate } from '@/helpers/validators';
import { API_SignUp } from '@/api/api';

export default{
  components:{
    loginInput,
    submitButton,
  },
  data(){
    return{
      name: {value: '', error: ''} as IValidObject,
      surname: {value: '', error: ''} as IValidObject,
      thirdname: {value: '', error: ''} as IValidObject,
      birthDate: {value: '', error: ''} as IValidObject,
      time: 4000,
    }
  },
  computed: {
    ...mapStores(useStatusWindowStore, useLoginStore),
  },
  methods:{
    SubmitSignUp(){
      if (this.name.value !== '' && this.surname.value !== '' && this.birthDate.value !== ''){
        const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Отправляем данные на сервер...', -1);
        const data: IApiSignUpData = {
          email: this.loginStore.SignupEmail,
          phone_number: this.loginStore.SignupPhoneNumber,
          password: this.loginStore.SignupPassword,
          name: this.loginStore.SignupName,
          surname: this.loginStore.SignupSurname,
          thirdname: this.loginStore.SignupThirdname,
          birth_date: this.loginStore.SignupBirthDate,
        };
        API_SignUp(data)
        .then(response => {
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Регистрация упешна!', this.time);
          this.$router.push({name: 'MainPage'});
        })
        .catch(error => {
          this.statusWindowStore.deteleStatusWindow(stID);
          this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Ошибка регистрации!', this.time);
        });

        return;
      }
      if(this.name.value === '') {
        if(this.name.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите ваше имя!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.name.error, this.time);
      }
      if(this.surname.value === '') {
        if(this.surname.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите вашу фамилию!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.surname.error, this.time);
      }
      if(this.thirdname.error !== '') {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, this.thirdname.error, this.time);
      }
      if(this.birthDate.value === '') {
        if(this.birthDate.error === '') this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Введите вашу дату рождения!', this.time);
        else this.statusWindowStore.showStatusWindow(StatusCodes.error, this.birthDate.error, this.time);
      }
    },
    ChangeName(value: string){
      this.name = ValidUserName(value); // валидируем
      if(this.name.error === '') this.loginStore.SignupName = this.name.value; // если корректно, то запоминаем
      if(value === '') this.name.error = ''; // если ничего не введено то убираем ошибку
    },
    ChangeSurname(value: string){
      this.surname = ValidUserSurname(value); // валидируем
      if(this.surname.error === '') this.loginStore.SignupSurname = this.surname.value; // если корректно, то запоминаем
      if(value === '') this.surname.error = ''; // если ничего не введено то убираем ошибку
    },
    ChangeThirdname(value: string){
      this.thirdname = ValidUserThirdname(value); // валидируем
      if(this.thirdname.error === '') this.loginStore.SignupThirdname = this.thirdname.value; // если корректно, то запоминаем
    },
    ChangeBirthDate(value: string){
      this.birthDate = ValidBirthDate(value); // валидируем
      if(this.birthDate.error === '') this.loginStore.SignupBirthDate= this.birthDate.value; // если корректно, то запоминаем
      if(value === '') this.birthDate.error = ''; // если ничего не введено то убираем ошибку
    }
  }
};
</script>