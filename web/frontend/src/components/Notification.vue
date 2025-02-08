<template>
    <div v-if="isVisible" class="modal-overlay">
      <div class="modal-content">
        <!-- Title -->
        <h3 class="modal-title">{{ title }}</h3>
  
        <!-- Message -->
        <p class="modal-message">{{ message }}</p>
  
        <!-- Buttons -->
        <div class="modal-buttons">
          <button class="modal-button" @click="handleCancel">{{ cancelText }}</button>
          <button class="modal-button primary-button" @click="handleConfirm">{{ confirmText }}</button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup lang="ts">
  import { defineProps, defineEmits } from "vue";
  
  // Define props for the modal
  defineProps({
    isVisible: {
      type: Boolean,
      required: true,
      default: false,
    },
    title: {
      type: String,
      required: true,
      default: "Notification",
    },
    message: {
      type: String,
      required: true,
      default: "This is a notification message.",
    },
    cancelText: {
      type: String,
      default: "Cancel",
    },
    confirmText: {
      type: String,
      default: "Continue",
    },
  });
  
  // Define emits for button actions
  const emit = defineEmits(["cancel", "confirm"]);
  
  // Handle cancel action
  const handleCancel = () => {
    emit("cancel");
  };
  
  // Handle confirm action
  const handleConfirm = () => {
    emit("confirm");
  };
  </script>
  
  <style scoped>
  /* Modal overlay */
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5); /* Semi-transparent background */
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 9999; /* Ensure it's on top */
  }
  
  /* Modal content */
  .modal-content {
    background-color: #ffffff;
    padding: 24px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    width: 90%;
    max-width: 400px;
    text-align: center;
  }
  
  /* Modal title */
  .modal-title {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 12px;
    color: #333;
  }
  
  /* Modal message */
  .modal-message {
    font-size: 14px;
    color: #666;
    margin-bottom: 20px;
  }
  
  /* Modal buttons */
  .modal-buttons {
    display: flex;
    gap: 10px;
    justify-content: center;
  }
  
  /* Button styles */
  .modal-button {
    padding: 10px 20px;
    font-size: 14px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .modal-button.primary-button {
    background-color: #007bff;
    color: #ffffff;
  }
  
  .modal-button.primary-button:hover {
    background-color: #0056b3;
  }
  
  .modal-button:not(.primary-button) {
    background-color: #ccc;
    color: #333;
  }
  
  .modal-button:not(.primary-button):hover {
    background-color: #aaa;
  }
  </style>