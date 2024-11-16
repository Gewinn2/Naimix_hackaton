//all constants

//api.ts
export const API = 'http://localhost:8080/';
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