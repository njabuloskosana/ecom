<template>
    <div class="tab-container">
      <button 
        class="tab-button" 
        :class="{ 'active': selectedTab === 'driver' }"
        @click="$emit('update:selectedTab', 'driver')"
      >
        I'm a Driver
      </button>
      <button 
        class="tab-button" 
        :class="{ 'active': selectedTab === 'bikeOwner' }"
        @click="$emit('update:selectedTab', 'bikeOwner')"
      >
        I'm a Bike Owner
      </button>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from 'vue';
  
  export default defineComponent({
    name: 'TabSelector',
    props: {
      selectedTab: {
        type: String,
        required: true
      }
    },
    emits: ['update:selectedTab']
  });
  </script>
  
  <style scoped>
  .tab-container {
    display: flex;
    width: 100%;
    max-width: 500px;
    margin: 0 auto 20px auto;
  }
  
  .tab-button {
    flex: 1;
    padding: 12px;
    text-align: center;
    background-color: #f5f7fa;
    border: none;
    cursor: pointer;
    transition: background-color 0.3s;
  }
  
  .tab-button.active {
    background-color: #e9eff9;
  }
  
  .tab-button:first-child {
    border-radius: 4px 0 0 4px;
  }
  
  .tab-button:last-child {
    border-radius: 0 4px 4px 0;
  }
  </style>