import { writable } from 'svelte/store';

export const user = writable(JSON.parse(localStorage.getItem('user')))

user.subscribe((value) => localStorage.user = JSON.stringify(value))