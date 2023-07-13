export interface Address {
  street: string;
  city: string;
  state: string;
  zipCode: string;
  CountryCode: string;
}

export interface User {
  firstName: string;
  lastName: string;
  email: string;
  address: Address;
}

export interface UserCredential {
  email: string;
  password: string;
}

export interface UserPostProps extends User, UserCredential {}
