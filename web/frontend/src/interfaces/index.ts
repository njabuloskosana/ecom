// Define the main response interface
export interface LoginResponse {
    jwtToken: string; // JWT token as a string
    userProfile: UserProfile; // User profile details
    userInformation: UserInformation; // User information details
  }
  
  // Define the UserProfile interface
  export interface UserProfile {
    id: number; // Unique identifier for the profile
    userId: number; // User ID associated with the profile
    firstName: string;
    lastName: string;
    phoneNumber: string;
    alternativeNumber?: string;
  }
  
  // Define the UserInformation interface
  export interface UserInformation {
    id: number; // Unique identifier for the user information
    uuId: string; // UUID of the user
    userEmail: string; // Email address of the user
    userName: string; // Username (can be the same as email)
    isActive: boolean; // Indicates if the user account is active
    creationDate: string; // Date when the user account was created
    userRoles: string[]; // Array of roles assigned to the user
    units: number; // Units associated with the user (e.g., loyalty points)
  }

export interface User {
        userEmail: string;
        userName: string;
        password: string;
        notificationPreferences: number[];
        roleId: number[];
        profile: UserProfile;
      }
      export      interface Profile {
        firstName: string;
        lastName: string;
        phoneNumber: string;
        alternativeNumber: string;
      }
      
      
export interface Store {
        id: number;
        name: string;
        address: string;
        hours: string;
        isOpen: boolean;
        hasLoyaltyProgram: boolean;
        comingSoon?: boolean;
      }
      
