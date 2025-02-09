<template>
    <div class="store-card">
      <h2>
        <img
          class="store-icon"
          :src="storeIcon"
          alt="Store Icon"
        />
        {{ store.name }}
      </h2>
      <p><strong>Address:</strong> {{ store.address }}</p>
      <p><strong>Hours:</strong> {{ store.hours }}</p>
      <p v-if="store.isOpen" class="status open">Status: Open Now</p>
      <p v-else class="status closed">Status: Closed</p>
      <p v-if="store.hasLoyaltyProgram">Rewards Program: Available</p>
      <button
        class="rewards-button"
        :disabled="store.comingSoon"
        @click="selectStore"
      >
        <template v-if="store.comingSoon">
          <i class="fas fa-ban disabled-fa"></i>
          Coming Soon
        </template>
        <template v-else>
          Claim Your Rewards Now
        </template>
      </button>
    </div>
  </template>
  
  <script lang="ts">
  import { defineComponent } from "vue";
  import type { PropType } from "vue";
  import storeIcon from "../assets/store.png";
  
  export interface Store {
    id: number;
    name: string;
    address: string;
    hours: string;
    isOpen: boolean;
    hasLoyaltyProgram: boolean;
    comingSoon?: boolean; // Optional: Indicates if the store is not yet available
  }
  
  export default defineComponent({
    name: "StoreCard",
    props: {
      store: {
        type: Object as PropType<Store>,
        required: true,
      },
    },
    emits: ["select-store"],
    setup(props, { emit }) {
      // Function to return the store icon path
  
      // Emits the "select-store" event when the store is clicked
      const selectStore = () => {
        if (!props.store.comingSoon) {
          emit("select-store", props.store);
        }
      };
  
      return { storeIcon, selectStore };
    },
  });
  </script>
  
  <style scoped>
    /* Store List */
    .store-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 20px;
  }

  .store-icon {
  width: 32px;
  height: 32px;
  margin-right: 8px;
  vertical-align: middle;
}
  
  /* Store Card */
  .store-card {
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 15px;
    background-color: var(--card-background);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
    width: 100%;
    height: 100%;
  }
  
  .store-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }
  
  .store-card h2 {
    margin-top: 0;
    margin-bottom: 10px;
    color: var(--primary-color);
  }
  
  .store-card p {
    margin: 5px 0;
    color: var(--secondary-color);
  }
  
  .store-card .status {
    font-weight: bold;
  }
  
  .store-card .status.open {
    color: green;
  }
  
  .store-card .status.closed {
    color: red;
  }
  
  .store-card button {
    margin-top: 10px;
    padding: 10px 20px;
    background-color: var(--primary-color);
    color: #fff;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
  }
  
  .store-card button:hover {
    background-color: var(--primary-hover);
  }
  
  /* No Results */
  .no-results {
    grid-column: 1 / -1;
    text-align: center;
    color: var(--secondary-color);
    font-size: 1.1rem;
  }
  
  /* Fade Transition */
  .fade-enter-active,
  .fade-leave-active {
    transition: opacity 0.3s ease;
  }
  
  .fade-enter-from,
  .fade-leave-to {
    opacity: 0;
  }

  .disabled-fa {
  margin-right: 5px;
  vertical-align: middle;
}

.rewards-button {
  padding: 10px 20px;
  background-color: var(--primary-color);
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.rewards-button:disabled {
  background-color: #aaa; /* Gray out disabled button */
  cursor: not-allowed;
}

.rewards-button:not(:disabled):hover {
  background-color: var(--primary-hover);
}
  </style>
  