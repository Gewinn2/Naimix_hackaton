//all constants

//api.ts
export const API = 'http://localhost:8080';
export const DEVMODE = true;
export const minDate = new Date('1950-01-01'); // минимальная дата рождения для пользователя
export const maxDate = new Date('2011-01-01'); // максимальная дата рождения для пользователя

//interfaces

export interface IValidObject{
  value: string,
  error: string,
};

export interface IStatusWindow{
  id: number,
  status: StatusCodes,
  text: string,
};

export interface IApiLoginData{
  email: string,
  phone_number: string,
  password: string,
};

export interface IApiSignUpData{
  email: string,
  phone_number: string,
  password: string,
  name: string,
  surname: string,
  thirdname: string,
  birth_date: string,
};

export interface IUsersList{
  id: number,
  name: string,
  role: string,
  isChoosed: boolean,
}

export interface IPickCard{
  isPicked: boolean,
};

export interface ICard{
  direct_meaning: string,
  reverse_meaning: string,
  id: number,
};

export interface ICardCategories{
  scope: string,
  category: string,
  description: string,
};

export interface ICosmogrammRequest{
  birth_date: string,
  name: string,
  surname: string,
  third_name: string,
};

export interface ICosmogrammResponse{
  compatibilities: ICosmogrammCompatibilitiy[],
  cosmogram: ICosmogrammCosmogramm,
};

export interface ICosmogrammCompatibilitiy{
  candidate_full_name: string,
  communication: number,
  communication_comment: string,
  emotions: number,
  emotions_comment: string,
  employee_full_name: string,
  work: number,
  work_comment: string,
};

export interface ICosmogrammCosmogramm{
  chiron: ICosmogrammPlanet,
  city: string,
  day: number,
  hour: number,
  jupiter: ICosmogrammPlanet,
  lat: number,
  lng: number,
  mars: ICosmogrammPlanet,
  mean_lilith: ICosmogrammPlanet,
  mean_node: ICosmogrammPlanet,
  mean_south_node: ICosmogrammPlanet,
  mercury: ICosmogrammPlanet,
  minute: number,
  month: number,
  moon: ICosmogrammPlanet,
  name: string,
  nation: string,
  neptune: ICosmogrammPlanet,
  pluto: ICosmogrammPlanet,
  saturn: ICosmogrammPlanet,
  sun: ICosmogrammPlanet,
  uranus: ICosmogrammPlanet,
  venus: ICosmogrammPlanet,
  year: number,
};

export interface ICosmogrammPlanet{
  abs_pos: number,
  name: string,
  retrograde: boolean,
};

//enums

export enum StatusCodes {
  'error', 'info', 'loading', 'success'
};

//cards

//all cards
export const TAROT_CARDS = [
  // Major Arcana
  'The_Fool',
  'The_Magician',
  'The_High_Priestess',
  'The_Empress',
  'The_Emperor',
  'The_Hierophant',
  'The_Lovers',
  'The_Chariot',
  'The_Power',
  'The_Hermit',
  'Wheel_of_Fortune',
  'Justice',
  'The_Hanged_Man',
  'Death',
  'Temperance',
  'The_Devil',
  'The_Tower',
  'The_Star',
  'The_Moon',
  'The_Sun',
  'Judgement',
  'The_World',

  // Minor Arcana - Wands
  'Ace of Wands',
  'Two of Wands',
  'Three of Wands',
  'Four of Wands',
  'Five of Wands',
  'Six of Wands',
  'Seven of Wands',
  'Eight of Wands',
  'Nine of Wands',
  'Ten of Wands',
  'Page of Wands',
  'Knight of Wands',
  'Queen of Wands',
  'King of Wands',

  // Minor Arcana - Cups
  'Ace of Cups',
  'Two of Cups',
  'Three of Cups',
  'Four of Cups',
  'Five of Cups',
  'Six of Cups',
  'Seven of Cups',
  'Eight of Cups',
  'Nine of Cups',
  'Ten of Cups',
  'Page of Cups',
  'Knight of Cups',
  'Queen of Cups',
  'King of Cups',

  // Minor Arcana - Swords
  'Ace of Swords',
  'Two of Swords',
  'Three of Swords',
  'Four of Swords',
  'Five of Swords',
  'Six of Swords',
  'Seven of Swords',
  'Eight of Swords',
  'Nine of Swords',
  'Ten of Swords',
  'Page of Swords',
  'Knight of Swords',
  'Queen of Swords',
  'King of Swords',

  // Minor Arcana - Pentacles
  'Ace of Pentacles',
  'Two of Pentacles',
  'Three of Pentacles',
  'Four of Pentacles',
  'Five of Pentacles',
  'Six of Pentacles',
  'Seven of Pentacles',
  'Eight of Pentacles',
  'Nine of Pentacles',
  'Ten of Pentacles',
  'Page of Pentacles',
  'Knight of Pentacles',
  'Queen of Pentacles',
  'King of Pentacles',
];

export function GET_CARD_IMAGE(cardID: number):string {
  const arcanaFolder = (cardID < 22) ? 'majorArcana' : 'minorArcana';
  return `${arcanaFolder}/${TAROT_CARDS[cardID]}.png`;
};

export const CARD_CATEGORIES = [
  {scope: 'Первая сфера:', category: 'Соответствие ожиданиям работодателя', description: 'Результат покажет, насколько навыки рекрута подходят под запрос работодателя'},
  {scope: 'Вторя сфера:', category: 'Мотивация и цели', description: 'Результат покажет, насколько карьерные цели и амбиции рекрута соотносятся с целями компании'},
  {scope: 'Третья сфера:', category: 'Финансовые перспективы', description: 'Результат покажет, какие финансовые условия и возможные доходы ждут рекрута'},
  {scope: 'Четвертая сфера:', category: 'Рабочая среда', description: 'Результат покажет, какая атмосфера и условия работы ждут рекрута'},
  {scope: 'Пятая сфера:', category: 'Личностный рост', description: 'Результат покажет, как рекрут сможет развиваться в выбранном направлении'},
  {scope: 'Шестая сфера:', category: 'Взаимоотношения с коллегами', description: 'Результат покажет, какие взаимоотношения ждут рекрута и будущих коллег'},
  {scope: 'Седьмая сфера:', category: 'Гибкость и адаптивность', description: 'Результат покажет, насколько быстро рекрут сможет адаптироваться к задачам'},
] as ICardCategories[];
