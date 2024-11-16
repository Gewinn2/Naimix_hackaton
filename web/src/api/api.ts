import axios from "axios";
import { API, DEVMODE, type IApiLoginData, type IApiSignUpData } from "@/helpers/constants";

//check authentication function
export function API_CheckAuth(){
  return new Promise((resolve, reject) => {
    axios.post(API)
    .then(response => {
      if(DEVMODE) console.log('CheckAuth success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('CheckAuth error: ', error);
      reject(error);
    });
  });
}

//login
export function API_Login(data:IApiLoginData){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/login`, data)
    .then(response => {
      if(DEVMODE) console.log('Login post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('Login post error: ', error);
      reject(error);
    });
  });
}

//signup
export function API_SignUp(data:IApiSignUpData){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/signup`, data)
    .then(response => {
      if(DEVMODE) console.log('SignUp post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('SignUp post error: ', error);
      reject(error);
    });
  });
}

//logout
export function API_LogOut(){
  return new Promise((resolve, reject) => {
    axios.post(`${API}/logout`)
    .then(response => {
      if(DEVMODE) console.log('LogOut post success: ', response);
      resolve(response);
    })
    .catch(error => {
      if(DEVMODE) console.log('LogOut post error: ', error);
      reject(error);
    });
  });
}