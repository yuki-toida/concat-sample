<template>
  <div>
    <h1>Video Concater</h1>
    <ul>
      <li><router-link v-bind:to="{ name: 'Home' }">Home</router-link></li>
    </ul>
    <hr>
    <router-view></router-view>
  </div>
</template>
