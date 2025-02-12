<template>
    <div class="icon-grid">
      <div v-for="(item, index) in items" :key="index" class="icon-item">
        <div class="icon-circle">
          <svg class="icon-svg" viewBox="0 0 24 24">
            <path :d="item.iconPath" fill="currentColor" />
          </svg>
        </div>
        <p class="icon-label">{{ item.label }}</p>
      </div>
    </div>
    </template>
    
    <script setup lang="ts">
    import { ref } from "vue";
    
    const items = ref([
    { label: "Receipt Photo", iconPath: "M3 2h13l5 5v14a2 2 0 01-2 2H3a2 2 0 01-2-2V4a2 2 0 012-2zm13 3.5V9h4.5L16 5.5zM6 8h6v2H6zm0 4h12v2H6zm0 4h12v2H6z" },
    { label: "Survey", iconPath: "M19 3h-4.18C14.4 1.84 13.3 1 12 1s-2.4.84-2.82 2H5a2 2 0 00-2 2v16a2 2 0 002 2h14a2 2 0 002-2V5a2 2 0 00-2-2zm-7-1a1 1 0 011 1h-2a1 1 0 011-1zm7 18H5V5h14z" },
    { label: "Scan QR Code", iconPath: "M4 4h4V2H2v6h2V4zm2 2H4v2h2V6zm4 4H6v2h4v-2zm6-6h4V2h-6v2h2v2zm-2 2h2v2h-2V6zm-4 6H6v2h4v-2zm4 4h2v2h-2v-2zm6 4h-4v2h6v-6h-2v4zM4 18H2v6h6v-2H4v-4zm4-2H6v2h2v-2zm10 0h-2v2h2v-2z" },
    { label: "Buy Points", iconPath: "M12 2L3.5 9H6v7h12V9h2.5zM12 5.5l4.5 3.5H7.5zM8 10h8v5H8z" },
    { label: "Invite Friends", iconPath: "M9 11a4 4 0 100-8 4 4 0 000 8zm6 0a4 4 0 100-8 4 4 0 000 8zM3 22v-2c0-2 4-3 6-3s6 1 6 3v2H3zm15 0v-2c0-1.2-1.6-2.2-3.5-2.7 1.7.5 3.5 1.5 3.5 2.7v2h2z" },
    { label: "Cashing Points", iconPath: "M12 2a10 10 0 1010 10A10 10 0 0012 2zm1 14h-2v-2h2zm0-4h-2V8h2z" },
    ]);
    </script>
    
    <style scoped>
    /* Grid Layout */
    .icon-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(120px, 1fr));
    gap: 20px;
    justify-items: center;
    margin: auto;
    padding: 20px;
    margin-bottom: 20px;
    }
    
    /* Individual Icon Item */
    .icon-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    transition: transform 0.2s ease-in-out;
    }
    
    .icon-item:hover {
    transform: scale(1.1);
    }
    
    /* Circular Icon Background */
    .icon-circle {
    width: 80px;
    height: 80px;
    background: linear-gradient(135deg, #1c86ee, #009acd);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s ease;
    }
    
    .icon-circle:hover {
    box-shadow: 0 6px 10px rgba(0, 0, 0, 0.2);
    }
    
    /* SVG Icon */
    .icon-svg {
    width: 40px;
    height: 40px;
    color: white;
    }
    
    /* Label Below Icons */
    .icon-label {
    margin-top: 10px;
    font-size: 14px;
    font-weight: 600;
    color: #333;
    }

    .icon-grid {
    border: 1px solid var(--border-color);
    border-radius: 8px;
    background-color: var(--card-background);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
  }
  
  .icon-grid:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
  
  
    
    /* Responsive Design */
    @media (max-width: 480px) {
    .icon-grid {
      grid-template-columns: repeat(2, 1fr);
    }
    
    .icon-circle {
      width: 70px;
      height: 70px;
    }
    
    .icon-svg {
      width: 35px;
      height: 35px;
    }
    
    .icon-label {
      font-size: 12px;
    }
    }
    
    @media (max-width: 768px) {
    .icon-grid {
      grid-template-columns: repeat(3, 1fr);
    }
    }
    </style>

