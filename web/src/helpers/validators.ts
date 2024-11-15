import { type IValidObject, minDate, maxDate } from "./constants";

//validate user name
export function ValidUserName(value: string): IValidObject {
  if(value.match(/^[a-zA-Zа-яА-Я]+$/) === null){
    return {value: '', error: 'Некорректное имя пользователя!'};
  }
  if(value.length < 2){
    return {value: '', error: 'Некорректное имя пользователя!'};
  }
  if(value.length > 20){
    return {value: '', error: 'Некорректное имя пользователя!'};
  }
  return {value: value, error: ''};
}

//validate user surname
export function ValidUserSurname(value: string): IValidObject {
  if(value.match(/^[a-zA-Zа-яА-Я]+[\-\s]?[a-zA-Zа-яА-Я]+$/) === null){
    return {value: '', error: 'Некорректная фамилия пользователя!'};
  }
  if(value.length < 3){
    return {value: '', error: 'Некорректная фамилия пользователя!'};
  }
  if(value.length > 25){
    return {value: '', error: 'Некорректная фамилия пользователя!'};
  }
  return {value: value, error: ''};
}

//validate user thirdname
export function ValidUserThirdname(value: string): IValidObject {
  if(value.length === 0){
    return {value: value, error: ''};
  }
  if(value.match(/^[a-zA-Zа-яА-Я]+[\-\s]?[a-zA-Zа-яА-Я]+$/) === null){
    return {value: '', error: 'Некорректное отчество пользователя!'};
  }
  if(value.length < 3){
    return {value: '', error: 'Некорректное отчество пользователя!'};
  }
  if(value.length > 25){
    return {value: '', error: 'Некорректное отчество пользователя!'};
  }
  return {value: value, error: ''};
}

//validate email
export function ValidEmail(value: string): IValidObject {
  if(value.match(/^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/) === null){
    return {value: '', error: 'Некорректный адрес электронной почты!'};
  }
  return {value: value, error: ''};
}

//validate phone number
export function ValidPhoneNumber(value: string): IValidObject {
  while(value.match(/[\-\s()]+/) !== null)value = value.replace(/[\-\s()]+/, ''); //+7 ( 999 ) - 888 - 77 66  ->  +79998887766
  if(value[0] === '+') value = value.slice(1); // +79998887766  ->  79998887766
  
  if(value.match(/^[\d]+$/) === null){
    return {value: '', error: 'Некорректный номер телефона!'};
  }
  if(value.length !== 11){
    return {value: '', error: 'Некорректный номер телефона!'};
  }
  return {value: value, error: ''};
}

//validate user password
export function ValidUserPassword(value: string): IValidObject {
  if(value.match(/[a-zA-Z]+/) === null){
    return {value: '', error: 'Пароль должен содержать латинские буквы в обоих регистрах!'};
  }
  if(value.match(/[a-z]+/) === null){
    return {value: '', error: 'Пароль должен содержать латинские буквы в нижнем регистре!'};
  }
  if(value.match(/[A-Z]+/) === null){
    return {value: '', error: 'Пароль должен содержать латинские буквы в верхнем регистре!'};
  }
  if(value.match(/[0-9]+/) === null){
    return {value: '', error: 'Пароль должен содержать хотя бы одну цифру!'};
  }
  if(value.match(/[!"№;%:\?\*()_\+`~@#\$\^&\-=]+/) === null){
    return {value: '', error: 'Пароль должен содержать хотя бы один спецсимвол!'};
  }
  if(value.match(/^[a-zA-Z0-9!"№;%:\?\*()_\+`~@#\$\^&\-=]+$/) === null){
    return {value: '', error: 'Некорректный пароль!'};
  }
  if(value.length < 6){
    return {value: '', error: 'Слишком короткий пароль!'};
  }
  if(value.length > 30){
    return {value: '', error: 'Слишком длинный пароль!'};
  }
  return {value: value, error: ''};
}

//validate birth date
export function ValidBirthDate(value: string): IValidObject {
  const date = new Date(value);
  if(date < minDate){
    return {value: '', error: 'Система не поддерживает пользователей, родившихся до 1950г!'};
  }
  if(date >= maxDate){
    return {value: '', error: 'Система не поддерживает пользователей, родившихся после 2011г!'};
  }
  return {value: value, error: ''};
}