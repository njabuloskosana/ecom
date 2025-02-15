<template>
    <div class="store-locator">
  
      <!-- Store Locator Title -->
      <Ads></Ads>
  
      <IconGrid></IconGrid>
      
      <!-- Search Bar -->
      <div class="search-bar">
        <input
          type="text"
          v-model="searchQuery"
          placeholder="Search for a store..."
          @keyup="filterStores"
        />
        <button @click="toggleFilters"><img src="../assets/filter.svg"></button>
      </div>
  
      <!-- Filters -->
      <transition name="fade">
        <div v-if="showFilters" class="filters">
          <label>
            <input type="checkbox" v-model="filters.openNow" @change="filterStores" /> Currently Open
          </label>
          <label>
            <input type="checkbox" v-model="filters.hasLoyaltyProgram" @change="filterStores" /> Has Loyalty Program
          </label>
        </div>
      </transition>

      <div class="label">
        <div>Favourites</div>
        <div class="see-all-label">See all</div>
      </div>
  
      <!-- Store List -->
      <div class="store-list">
        <div v-if="filteredStores.length === 0" class="no-results">
          No stores found matching your criteria.
        </div>
        <div  v-for="store in filteredStores">
          <StoreCard
      :key="store.id"
      :store="store"
      @select-store="selectStore"
    />
        </div>
        <!-- Gaming Center -->
        <div class="store-card">
          <h2>
          <img
            class="store-icon"
            :src="gamingIcon"
          />
          {{ `Gaming Center` }}
          </h2>
          <p><strong>Status:</strong> {{ `Online` }}</p>
          <p><strong>Hours:</strong> {{ `24 Hours` }}</p>
          <p> <strong>Gaming Center:</strong>Unleash your skills, conquer challenges, and win big!</p>
          <p class="status open">Status: Open Now</p>
          <button
        class="rewards-button"
        @click="openGamingCenter()"
        >
        Start Playing
        </button>
        </div>

        <!-- Customer Service -->
        <div class="store-card">
          <h2>
          <img
            class="store-icon"
            :src="supportIcon"
          />
          {{ `Customer Service Center` }}
          </h2>
          <p><strong>Status:</strong> {{ `Online` }}</p>
          <p><strong>Hours:</strong> {{ `24 Hours` }}</p>
          <p> <strong>Customer Service:</strong> Helping with offers, rewards , promotions, questions, and account support.</p>
          <p class="status open">Status: Open Now</p>
          <button
        class="rewards-button"
        @click="contactSupport()"
        >
        Contact Support
        </button>
        </div>

        <!-- Rewards Network -->
        <div class="store-card">
        <h2>
          <img class="store-icon" :src="registerIcon" />
          {{ `Join SA's Best Rewards Network ` }}
        </h2>
        <p><strong>Status:</strong> {{ `Online` }}</p>
        <p><strong>Hours:</strong> {{ `24 Hours` }}</p>
        <p><strong>Register your business today , </strong> Partner with us and unlock exclusive benefits. </p>
        <p class="status open">Status: Available</p>

        <button class="rewards-button" @click="registerForRewards()">
          Join the Rewards Network
        </button>

        </div>

        <!-- Shop for Rewards -->
        <div class="store-card">
        <h2>
          <img class="store-icon" :src="shopIcon" />
          {{ `Shop for Loyalty Points and Rewards ` }}
        </h2>
        <p><strong>Status:</strong> {{ `Online` }}</p>
        <p><strong>Hours:</strong> {{ `24 Hours` }}</p>
        <p><strong>Shop for loyalty points and rewards,</strong> Earn points with every purchase and unlock exclusive benefits.</p>
        <p class="status open">Status: Available</p>

        <button class="rewards-button" @click="registerForRewards()">
          Shop For Rewards
        </button>

        </div>
      </div>
    </div>
    <BottomNav
      :navItems="navItems"
      :currentRoute="'/'"
    />
  </template>
  
  <script lang="ts">
  import supportIcon from "../assets/support.png";
  import gamingIcon from "../assets/gaming.png";
  import shopIcon from "../assets/discount.png";
  import registerIcon from "../assets/register.png";
  import discoverIcon from "../assets/discover.svg";
import storeIcon from "../assets/cart.svg";
import redeemIcon from "../assets/loyalty_program.svg";
import walletIcon from "../assets/wallet.svg";
import { defineComponent, ref, computed } from "vue";
import StoreCard from "../components/StoreCard.vue";
import { useRouter } from 'vue-router';
import type { Store } from "@/interfaces";
import IconGrid from "@/components/IconGrid.vue";
import BottomNav from "../components/BottomNav.vue";
import Ads from "../components/Ads.vue"
  
  export default defineComponent({
    name: "StoreLocator",
    components: {
      StoreCard,
      IconGrid,
      BottomNav,
      Ads,
    },
    setup() {

 const navItems = [
        {
          label: 'Discover',
          route: '/',
          svgSrc: discoverIcon
        },
        {
          label: 'Shop',
          route: '/coming-soon',
          svgSrc: storeIcon
        },
        {
          label: 'Redeem',
          route: '/coming-soon',
          svgSrc: redeemIcon
        },
        {
          label: 'Wallet',
          route: '/coming-soon',
          svgSrc: walletIcon
        }
      ]
      const router = useRouter();
      // Sample store data
      const stores = ref<Store[]>([
        {
          id: 1,
          name: "Total Elardus Park",
          address: "Solomon Mahlangu Dr, Erasmuskloof, Pretoria, 0181",
          hours: "24 Hours",
          isOpen: true,
          hasLoyaltyProgram: true,
          comingSoon: false,
        },
        /*{
          id: 2,
          name: "Timetreasure Laundry",
          address: "14 Panfluit Eco-Park East, Centurion, 0144",
          hours: "24 Hours",
          isOpen: false,
          hasLoyaltyProgram: false,
          comingSoon: true,
        }*/
      ]);
  
      const searchQuery = ref<string>("");
      const showFilters = ref<boolean>(false);
      const filters = ref({
        openNow: false,
        hasLoyaltyProgram: false,
      });
  
      // Filtered stores based on search and filters
      const filteredStores = computed(() => {
        return stores.value.filter((store) => {
          const matchesSearch =
            store.name.toLowerCase().includes(searchQuery.value.toLowerCase()) ||
            store.address.toLowerCase().includes(searchQuery.value.toLowerCase());
          const matchesFilters =
            (!filters.value.openNow || store.isOpen) &&
            (!filters.value.hasLoyaltyProgram || store.hasLoyaltyProgram);
          return matchesSearch && matchesFilters;
        });
      });
  
      const toggleFilters = () => {
        showFilters.value = !showFilters.value;
      };
  
      const filterStores = () => {
        // This function triggers reactivity when filters are updated
      };
  
      const selectStore = (store: Store) => {
        router.push('/login');
        // Implement navigation or further actions here
      };

       const contactSupport = () => {
        router.push('/coming-soon');
        // Implement navigation or further actions here
      };

      const openGamingCenter = () => {
        router.push('/games/tictactoe');
        // Implement navigation or further actions here
      };

      const registerForRewards = () => {
        router.push('/coming-soon');
        // Implement navigation or further actions here
      };
  
      return {
        stores,
        searchQuery,
        showFilters,
        filters,
        filteredStores,
        toggleFilters,
        filterStores,
        selectStore,
        contactSupport,
        supportIcon,
        openGamingCenter,
        gamingIcon,
        registerForRewards,
        registerIcon,
        redeemIcon,
        walletIcon,
        shopIcon,
        navItems,
        discoverIcon,
        storeIcon,

      };
    },
  });
  </script>
  
  <style scoped>
  .store-icon {
  width: 32px;
  height: 32px;
  margin-right: 8px;
  vertical-align: middle;
}
  /* Container */
  .store-locator {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
    background-color: var(--background-color);
    min-height: 100vh;
    font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
  }
  
  /* Welcome Screen */
  .welcome-screen {
    text-align: center;
    margin-bottom: 20px;
    padding: 20px;
    border: 1px solid var(--border-color);
    border-radius: 8px;
    background-color: var(--card-background);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
  }
  
  /* Locator Title */
  .locator-title {
    text-align: center;
    margin-bottom: 20px;
    color: var(--secondary-color);
  }
  
  .search-bar {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 12px;
  margin: 1.25rem 0;
  padding: 0.5rem;
  background-color: var(--secondary-bg);
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}

.search-bar input {
  flex: 1;
  padding: 0.75rem 1rem;
  border: 1px solid var(--border-color);
  border-radius: 6px;
  font-size: 1rem;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
}

.search-bar input:focus {
  outline: none;
  border-color: var(--primary-color);
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.search-bar button {
  padding: 0.75rem 1.5rem;
  background-color: var(#888);
  color: #fff;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 1rem;
  font-weight: 500;
  transition: background-color 0.3s ease, transform 0.2s ease;
}

.search-bar button:hover {
  background-color: var(#888);
  transform: translateY(-1px);
}
  /* Filters */
  .filters {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .filters label {
    margin-right: 15px;
    font-size: 0.9rem;
    color: var(--secondary-color);
  }
  
  /* Store List */
  .store-list {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 20px;
  }
  
  /* Store Card */
  .store-card {
    border: 1px solid var(--border-color);
    border-radius: 8px;
    padding: 15px;
    background-color: var(--card-background);
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.05);
    transition: transform 0.2s ease, box-shadow 0.2s ease;
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

.label{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  width: 100%;
  margin-bottom: 10px;
  color: var(--secondary-color);
  font-size: 20px;
  padding: 5px;
}

.see-all-label{
  color:black;
}
  </style>
  