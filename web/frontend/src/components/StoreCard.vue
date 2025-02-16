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
      <p v-if="store.isOpen" class="status open">Available</p>
      <p v-else class="status closed">Status: Closed</p>
      <p v-if="store.hasLoyaltyProgram">Rewards Program</p>
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
    .store-icon {
  width: 32px;
  height: 32px;
  margin-right: 8px;
  vertical-align: middle;
}

    .store-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 20px;
  }
  
  /* Store Card */
  .store-card {
  border: 1px solid var(--border-color, #e0e0e0);
  border-radius: 10px;
  padding: 20px;
  background-color: var(--card-background, #fff);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease, box-shadow 0.3s ease;
  display: flex;
  flex-direction: column;
 
}

.store-card:hover {
  transform: translateY(-5px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.store-card h2 {
  margin: 0 0 12px;
  color: #3c336b;
  font-size: 1.5rem;
  font-weight: 600;
}

.store-card p {
  display: inline-block;
  padding: 6px 12px;
  border-radius: 16px;
  margin: 4px 4px 0 0;
  font-size: 0.875rem;
  background-color: #f0f0f0;
  color: var(--secondary-color, #555);
  line-height: 1.5;
}

/* Specific chip style for status elements */
.store-card .status {
  display: inline-block;
  padding: 6px 12px;
  border-radius: 16px;
  font-weight: bold;
  margin: 10px 4px 10px 0;
  font-size: 0.875rem;
  background-color: #f0f0f0;
}

/* Open status chip */
.store-card .status.open {
  background-color: #e6f9ee; /* light green background */
  color: var(--success-color, #28a745);
}

/* Closed status chip */
.store-card .status.closed {
  background-color: #ffe6e6; /* light red background */
  color: var(--danger-color, #dc3545);
}

.store-card button {
  align-self: flex-start;
  margin-top: 10px;
  padding: 10px 25px;
  background-color: #3c336b;
  color: #fff;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: background-color 0.3s ease, transform 0.2s ease;
}

.store-card button:hover {
  background-color: var(--primary-hover, #0056b3);
  transform: scale(1.02);
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
  