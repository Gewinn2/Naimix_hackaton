<template>
  <div class="w-full h-full flex flex-col items-center overflow-y-scroll">
    <div class="w-8/12 h-full flex flex-col z-20 relative top-5">
      <div class="absolute z-10 right-0 flex flex-col justify-center items-center">
        <svg class="absolute w-[500px] h-[300px] z-10" viewBox="0 0 766 404" fill="none" xmlns="http://www.w3.org/2000/svg">
          <circle cx="250" cy="154" r="250" fill="#9A9A9A"/>
          <circle cx="590" cy="154" r="250" fill="#FF9929"/>
        </svg>
        <p class="text-2xl font-medium text-center z-20 relative">Воспользуйся нашей космограммой,<br> чтобы провести анализ совместимости<br> рекрута и будущих коллег.</p>
      </div>
      <section>
        <div class="flex flex-row items-center self-start w-full mt-5 mb-10 gap-x-2">
          <svg class="w-16 h-16" viewBox="0 0 60 70" fill="none" xmlns="http://www.w3.org/2000/svg">
            <g clip-path="url(#clip0_77_379)">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M4.8777 1.90252C6.3441 -0.265109 9.03904 -0.635434 10.897 1.07537C15.2899 5.12036 18.8019 10.3185 21.1471 16.2466C22.8684 20.5978 23.9254 25.252 24.2823 30.0002H35.7182C36.075 25.252 37.132 20.5978 38.8533 16.2466C41.1985 10.3185 44.7107 5.12036 49.1036 1.07537C50.9614 -0.635434 53.6563 -0.265109 55.1229 1.90252C56.589 4.07014 56.2719 7.21426 54.4136 8.92506C51.0806 11.9944 48.4153 15.9388 46.6359 20.4371C45.4354 23.4717 44.6606 26.7004 44.3331 30.0002H55.7143C58.0813 30.0002 60 32.2388 60 35.0002C60 37.7616 58.0813 40.0002 55.7143 40.0002H44.3331C44.6606 43.3001 45.4354 46.5289 46.6359 49.5635C48.4153 54.0617 51.0806 58.0062 54.4136 61.0757C56.2719 62.7862 56.589 65.9302 55.1229 68.0982C53.6563 70.2657 50.9614 70.6362 49.1036 68.9252C44.7107 64.8802 41.1985 59.6822 38.8533 53.7542C37.132 49.4028 36.075 44.7485 35.7182 40.0002H24.2823C23.9255 44.7485 22.8685 49.4028 21.1471 53.7542C18.8019 59.6822 15.2899 64.8802 10.897 68.9252C9.03904 70.6362 6.3441 70.2657 4.8777 68.0982C3.41129 65.9302 3.72871 62.7862 5.58669 61.0757C8.92007 58.0062 11.585 54.0617 13.3646 49.5635C14.5651 46.5289 15.3397 43.3001 15.6675 40.0002H4.28571C1.91878 40.0002 0 37.7616 0 35.0002C0 32.2388 1.91878 30.0002 4.28571 30.0002H15.6674C15.3397 26.7004 14.565 23.4717 13.3646 20.4371C11.585 15.9388 8.92007 11.9944 5.58669 8.92506C3.72871 7.21426 3.41129 4.07014 4.8777 1.90252Z" fill="white"/>
            </g>
            <defs>
            <clipPath id="clip0_77_379">
            <rect width="60" height="70" fill="white"/>
            </clipPath>
            </defs>
          </svg>
          <p class="text-4xl font-semibold">Космограмма</p>
        </div>
        <UserList 
          class="w-full mb-2"
          title="Выберите от 1 до 10 коллег для анализа совместимости" 
          placeholder="Введите ФИО" 
          :item-component="getUserItem"/>
      </section>
      <section>
        <div class="flex flex-row gap-x-2 text-lg items-center py-2">
          <p>Введите год рождения и ФИО рекрута: </p>
          <div>
            <input 
              type="date" 
              v-model="recrutBirthDay"
              class="outline-none py-2 px-4 bg-slate-50 text-lg text-gray-800 font-medium border-2 rounded-lg border-solid border-gray-400"/>
          </div>
          <div>
            <input 
              type="text" 
              v-model="recrutFIO"
              class="outline-none py-2 px-4 bg-slate-50 text-lg text-gray-800 font-medium border-2 rounded-lg border-solid border-gray-400 max-w-none w-[350px]" placeholder="ФИО рекрута"/>
          </div>
          <div v-if="1 <= choosedUsers.length && choosedUsers.length <= 10" @click="analyseCosmogramm" class="flex flex-row items-center justify-center py-2 px-4 rounded-lg cursor-pointer btn text-slate-50 text-gl font-medium">
            Анализировать
          </div>
        </div>
      </section>
      <section v-if="getCosmogrammPlanets.length !== 0">
        <div class="flex flex-row w-full justify-between gap-x-4">
          <svg id="cosmogrammSVG"></svg>
          <div class="flex flex-col items-center flex-grow gap-y-10">
            <p class="text-center text-3xl font-medium">Данные космограммы</p>
            <div class="flex flex-col gap-y-4 items-start w-full">
              <p class="text-lg cursor-default">Солнце: <span class="cursor-pointer">{{ getSun }}&#176;</span></p>
              <p class="text-lg cursor-default">Меркурий: <span class="cursor-pointer">{{ getMercury }}&#176;</span></p>
              <p class="text-lg cursor-default">Венера: <span class="cursor-pointer">{{ getVenus }}&#176;</span></p>
              <p class="text-lg cursor-default">Луна: <span class="cursor-pointer">{{ getMoon }}&#176;</span></p>
              <p class="text-lg cursor-default">Марс: <span class="cursor-pointer">{{ getMars }}&#176;</span></p>
              <p class="text-lg cursor-default">Юпитер: <span class="cursor-pointer">{{ getJupiter }}&#176;</span></p>
              <p class="text-lg cursor-default">Сатурн: <span class="cursor-pointer">{{ getSaturn }}&#176;</span></p>
              <p class="text-lg cursor-default">Уран: <span class="cursor-pointer">{{ getUranus }}&#176;</span></p>
              <p class="text-lg cursor-default">Нептун: <span class="cursor-pointer">{{ getNeptune }}&#176;</span></p>
              <p class="text-lg cursor-default">Плутон: <span class="cursor-pointer">{{ getPluto }}&#176;</span></p>
              <p class="text-lg cursor-default">Лилит: <span class="cursor-pointer">{{ getLilith }}&#176;</span></p>
              <p class="text-lg cursor-default">Нисх. узел: <span class="cursor-pointer">{{ getMSN}}&#176;</span></p>
              <p class="text-lg cursor-default">Восх. узел: <span class="cursor-pointer">{{ getMN }}&#176;</span></p>
            </div>
          </div>
        </div>
      </section>
      <section>
        <p class="text-4xl text-center">Анализ совместимости</p>
        <div class="flex flex-col gap-y-8">
          <div class="flex flex-col gap-y-4">
            <p class="text-2xl">Коммуникация: 100%</p>
            <p class="text-lg text-justify">A customer calls, extremely upset about a recent order error. They are speaking rapidly and interrupting you. Describe your approach to de-escalating the situation, actively listening to their concerns, and resolving the problem to the customer’s satisfaction. What specific communication strategies would you use?</p>
          </div>
          <div class="flex flex-col gap-y-4">
            <p class="text-2xl">Эмоциональность: 67%</p>
            <p class="text-lg text-justify">A customer calls, extremely upset about a recent order error. They are speaking rapidly and interrupting you. Describe your approach to de-escalating the situation, actively listening to their concerns, and resolving the problem to the customer’s satisfaction. What specific communication strategies would you use?</p>
          </div>
          <div class="flex flex-col gap-y-4">
            <p class="text-2xl">Работоспособность: 93%</p>
            <p class="text-lg text-justify">A customer calls, extremely upset about a recent order error. They are speaking rapidly and interrupting you. Describe your approach to de-escalating the situation, actively listening to their concerns, and resolving the problem to the customer’s satisfaction. What specific communication strategies would you use?</p>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>
<script lang="ts">
import { mapStores } from 'pinia';
import { useCosmogrammStore } from '@/stores/cosmogrammStore';
import { useStatusWindowStore } from '@/stores/statusWindowStore';
import UserList from '@/entities/userList.vue';
import { type IUsersList, type ICosmogrammRequest, minDate, maxDate, StatusCodes, type ICosmogrammPlanet } from '@/helpers/constants';
import UserItem from '@/shared/userItem.vue';
import { API_Cosmogramm } from '@/api/api';
import { generateCosmogramm, calculatePointCoordinates, drawLine, type Point, checkAngleDifference, type ILineType} from '@/helpers/cosmogrammDrawer';
export default{
  components:{
    UserList,
    UserItem,
  },
  data(){
    return{
      recrutBirthDay: '',
      recrutFIO: '',
    }
  },
  computed:{
    ...mapStores(useCosmogrammStore, useStatusWindowStore),

    getUserItem(){
      return UserItem;
    },

    choosedUsers(){
      return this.cosmogrammStore.usersList.filter(item => item.isChoosed);
    },

    getCosmogrammPlanets(){
      return this.cosmogrammStore.cosmoPlanets;
    },

    getSun() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Sun')[0]?.abs_pos.toFixed(2); },
    getMercury() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Mercury')[0]?.abs_pos.toFixed(2); },
    getVenus() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Venus')[0]?.abs_pos.toFixed(2); },
    getMoon() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Moon')[0]?.abs_pos.toFixed(2); },
    getMars() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Mars')[0]?.abs_pos.toFixed(2); },
    getJupiter() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Jupiter')[0]?.abs_pos.toFixed(2); },
    getSaturn() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Saturn')[0]?.abs_pos.toFixed(2); },
    getUranus() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Uranus')[0]?.abs_pos.toFixed(2); },
    getNeptune() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Neptune')[0]?.abs_pos.toFixed(2); },
    getPluto() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Pluto')[0]?.abs_pos.toFixed(2); },
    getLilith() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Mean_Lilith')[0]?.abs_pos.toFixed(2); },
    getMSN() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Mean_South_Node')[0]?.abs_pos.toFixed(2); },
    getMN() { return this.cosmogrammStore.cosmoPlanets.filter(i => i.name === 'Mean_Node')[0]?.abs_pos.toFixed(2); },
  },
  mounted(){
    if(this.cosmogrammStore.cosmoPlanets.length !== 0) this.drawCosmogram();
  },
  methods: {
    analyseCosmogramm(){
      const choosedUser = this.cosmogrammStore.usersList.filter(item => item.isChoosed);
      if(choosedUser.length < 1 || choosedUser.length > 10) return;
      const recrutBD = new Date(this.recrutBirthDay);
      if(recrutBD < minDate){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Система не поддерживает пользователей, родившихся до 1950г!', 3000);
        return;
      }
      if(recrutBD > maxDate){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Система не поддерживает пользователей, родившихся после 2011г!', 3000);
        return;
      }
      const recrutFIO = this.recrutFIO.split(' '); 
      if(recrutFIO.length < 2 || recrutFIO.length > 3){
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Некорректное ФИО рекрута!', 3000);
        return;
      }

      const data: ICosmogrammRequest = {  
        birth_date: this.recrutBirthDay,
        name: recrutFIO[0],
        surname: recrutFIO[1],
        third_name: recrutFIO[2] ? recrutFIO[2] : '',
      };

      const stID = this.statusWindowStore.showStatusWindow(StatusCodes.loading, 'Создаем космограмму совместимости...', -1);
      API_Cosmogramm(data)
      .then(res => {
        this.statusWindowStore.showStatusWindow(StatusCodes.success, 'Космограмма готова!', 3000);
        //
      })
      .catch(error => {
        this.statusWindowStore.showStatusWindow(StatusCodes.error, 'Что-то пошло не так при создании космограммы!', 3000);
        //
      })
      .finally(() => {
        this.statusWindowStore.deteleStatusWindow(stID);

        const response = {
          compatibilities: [
            {
              "candidate_full_name": "Doe John Petrovich",
              "employee_full_name": "Doe John Petrovich",
              "communication": 100,
              "communication_comment": "Ваши способы общения идеально совпадают. Вы понимаете друг друга с полуслова, и ваши разговоры всегда наполнены смыслом и взаимопониманием. Это создает крепкую основу для доверительных отношений.",
              "emotions": 100,
              "emotions_comment": "Вы делитесь глубокими эмоциональными связями, которые позволяют вам чувствовать себя комфортно и безопасно рядом друг с другом. Это создает прочный фундамент для ваших отношений, полных любви и поддержки.",
              "work": 0,
              "work_comment": "Ваши профессиональные подходы кажутся абсолютно несовместимыми. Это может создавать значительные препятствия в работе вместе и мешать достижению общих целей. Возможно, стоит пересмотреть свои ожидания и обсудить, как вы можете лучше сотрудничать."
            }
          ],
          cosmogram: {
            "name": "Doe John Petrovich",
            "year": 1990,
            "month": 1,
            "day": 1,
            "hour": 12,
            "minute": 11,
            "city": "Moscow",
            "nation": "RU",
            "lng": 37.6156,
            "lat": 55.7522,
            "sun": {
              "name": "Sun",
              "abs_pos": 280.6939374005611,
              "retrograde": false
            },
            "moon": {
              "name": "Moon",
              "abs_pos": 331.67934129539293,
              "retrograde": false
            },
            "mercury": {
              "name": "Mercury",
              "abs_pos": 295.70619328713667,
              "retrograde": true
            },
            "venus": {
              "name": "Venus",
              "abs_pos": 306.2369170971076,
              "retrograde": true
            },
            "mars": {
              "name": "Mars",
              "abs_pos": 249.9170335186626,
              "retrograde": false
            },
            "jupiter": {
              "name": "Jupiter",
              "abs_pos": 95.16460072135177,
              "retrograde": true
            },
            "saturn": {
              "name": "Saturn",
              "abs_pos": 285.6435644065208,
              "retrograde": false
            },
            "uranus": {
              "name": "Uranus",
              "abs_pos": 275.778392524542,
              "retrograde": false
            },
            "neptune": {
              "name": "Neptune",
              "abs_pos": 282.03364522650554,
              "retrograde": false
            },
            "pluto": {
              "name": "Pluto",
              "abs_pos": 227.08993880185324,
              "retrograde": false
            },
            "chiron": {
              "name": "Chiron",
              "abs_pos": 103.8213431856962,
              "retrograde": true
            },
            "mean_lilith": {
              "name": "Mean_Lilith",
              "abs_pos": 216.45077052953468,
              "retrograde": false
            },
            "mean_node": {
              "name": "Mean_Node",
              "abs_pos": 318.44123384913485,
              "retrograde": true
            },
            "mean_south_node": {
              "name": "Mean_South_Node",
              "abs_pos": 138.44123384913485,
              "retrograde": true
            }
          }
        };

        this.cosmogrammStore.cosmoPlanets = [];
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.chiron);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.jupiter);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.mars);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.mean_lilith);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.mean_node);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.mean_south_node);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.mercury);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.moon);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.neptune);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.pluto);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.saturn);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.sun);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.uranus);
        this.cosmogrammStore.cosmoPlanets.push(response.cosmogram.venus);

        setTimeout(() => {
          this.drawCosmogram(),
          200
        });
      });
    },
    drawCosmogram(){
      const container = document.getElementById("cosmogrammSVG");
      if(container === null) return;
      const width = 800;
      container.setAttribute("width", width.toString());
      container.setAttribute("height", width.toString());
      const radius = generateCosmogramm(container, width);

      for(let angle1 of this.cosmogrammStore.cosmoPlanets){
        for(let angle2 of this.cosmogrammStore.cosmoPlanets){
          if(angle1 === angle2) continue;

          const lineType = checkAngleDifference(angle1.abs_pos, angle2.abs_pos);
          if(!lineType.show) continue;

          drawLine(
            container, 
            calculatePointCoordinates(width / 2, width / 2, radius, angle1.abs_pos),
            calculatePointCoordinates(width / 2, width / 2, radius, angle2.abs_pos),
            ['line', lineType.class]
          );
        }
      }
    },
    getPlanet(name: string): ICosmogrammPlanet{
      this.cosmogrammStore.cosmoPlanets.forEach(item => {
        if(item.name === name) return item;
      });
      return {abs_pos: 0, name: '', retrograde: false};
    }
  }
};
</script>