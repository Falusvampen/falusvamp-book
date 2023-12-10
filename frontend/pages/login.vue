<template>
  <div class="container">
    <div class="form-container">
      <!-- <transition name="fade"> -->
      <form
        v-if="!showRegistrationForm"
        @submit.prevent="submitLoginForm"
        key="loginForm"
        class="form"
      >
        <h2>Welcome to FalusvampenBook</h2>
        <div class="input-group">
          <label for="login-email">Email:</label>
          <input type="email" id="login-email" v-model="email" required />
        </div>
        <div class="input-group">
          <label for="login-password">Password:</label>
          <input
            type="password"
            id="login-password"
            v-model="password"
            required
          />
        </div>
        <button class="btn-submit">Login</button>
        <button class="btn-switch" @click="showRegistrationForm = true">
          Not a user? Register here
        </button>
      </form>

      <form
        v-else
        @submit.prevent="submitRegistrationForm"
        key="registerForm"
        class="form"
      >
        <h2>Join our cult today!</h2>
        <div class="input-group">
          <label for="nickname">Nickname:</label>
          <input type="text" id="nickname" v-model="nickname" />
        </div>
        <div class="input-group">
          <label for="register-email">Email:</label>
          <input type="email" id="register-email" v-model="email" required />
        </div>
        <div class="input-group">
          <label for="register-password">Password:</label>
          <input
            type="password"
            id="register-password"
            v-model="password"
            required
          />
        </div>
        <div class="input-group grid-2">
          <div>
            <label for="first-name">First Name:</label>
            <input type="text" id="first-name" v-model="firstName" required />
          </div>
          <div>
            <label for="last-name">Last Name:</label>
            <input type="text" id="last-name" v-model="lastName" required />
          </div>
        </div>
        <div class="input-group grid-2">
          <div>
            <label for="date-of-birth">Date of Birth:</label>
            <input
              type="date"
              id="date-of-birth"
              v-model="dateOfBirth"
              required
            />
          </div>
          <div>
            <label for="avatar">Avatar/Image:</label>
            <input
              type="file"
              id="avatar"
              accept="image/*"
              @change="handleAvatarChange"
            />
          </div>
        </div>
        <div class="input-group">
          <label for="about-me">About Me:</label>
          <textarea id="about-me" v-model="aboutMe"></textarea>
        </div>
        <button class="btn-submit">Register</button>
        <button class="btn-switch" @click="showRegistrationForm = false">
          Already a user? Login here
        </button>
      </form>
      <!-- </transition> -->
    </div>
  </div>
</template>

<script>
// import { resolveDirective } from "nuxt/dist/app/compat/capi";

// const { $getEmailFromCookie } = useNuxtApp();

import { useAuthUser } from '/store/authUser';

export default {
  data() {
    return {
      showRegistrationForm: false,
      email: "",
      password: "",
      firstName: "",
      lastName: "",
      dateOfBirth: "",
      // avatar: null,
      nickname: "",
      aboutMe: "",
    };
  },
  methods: {
    submitLoginForm() {
      const authUser = useAuthUser(); // Create an instance of the store
      // Retrieve form data and perform login logic here
      const formData = {
        email: this.email,
        password: this.password,
      };

      // Example: Send the form data to a login API endpoint
      fetch("http://localhost:8080/api/user/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: 'include',
        body: JSON.stringify(formData),
      })
        .then((response) => {
          if (response.ok) {
            response.json().then(async (data) => {
              document.cookie = `user_uuid=${data.UUID}; path=/`;
              if (data.UUID) {
                await authUser.setAuthUser(true, data.UUID);
                // authUser.setAuthUser(true, data.UUID);  // Call the action on the store instance
                navigateTo("/");
              } else {
                throw new Error("User uuid not set");
              }
            });
          } else {
            // Login failed, handle the error
            console.log("Login failed");
            response.json().then((error) => {
              console.log(error);
            });
          }
        })
        .catch((error) => {
          console.error("Error occurred while logging in:", error);
          // Handle network or fetch error
        });
    },
    submitRegistrationForm() {
      const requestData = {
        email: this.email,
        firstName: this.firstName,
        lastName: this.lastName,
        password: this.password,
        dob: this.dateOfBirth,
        avatar: "",
        nickname: this.nickname,
        aboutMe: this.aboutMe,
        public: true, // Assuming the "public" field is always set to true
      };

      console.log(requestData);

      if (this.avatar) {
        // Convert the avatar image to base64 encoded string
        const reader = new FileReader();
        reader.onloadend = () => {
          requestData.avatar = reader.result;
          this.sendRegistrationRequest(requestData);
        };
        reader.readAsDataURL(this.avatar);
      } else {
        this.sendRegistrationRequest(requestData);
      }
    },
    sendRegistrationRequest(requestData) {
      fetch("http://localhost:8080/api/user/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestData),
      })
        .then((response) => {
          if (response.ok) {
            // Registration successful, redirect the user or perform other actions
            console.log("Registration successful");
            this.showRegistrationForm = false;
            navigateTo("/");
          } else {
            // Registration failed, handle the error
            console.log("Registration failed");
            response.json().then((error) => {
              console.log(error);
            });
          }
        })
        .catch((error) => {
          console.error("I am error-ful", error, error.error, error.message);
          // Handle network or fetch error
        });
    },
  },
  handleAvatarChange(event) {
    this.avatar = event.target.files[0];
  },
};
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;600&display=swap");

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f8f9fa;
  font-family: "Poppins", sans-serif;
}

.form-container {
  background: white;
  padding: 2rem;
  border-radius: 0.5rem;
  width: 100%;
  max-width: 600px;
}

.form h2 {
  margin-bottom: 1.5rem;
  color: #333;
  text-align: center;
}

.input-group {
  margin-bottom: 1.5rem;
}

.input-group label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 600;
  color: #333;
}

.input-group input,
.input-group textarea {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 0.5rem;
  font-size: 1rem;
  line-height: 1.5;
}

.input-group textarea {
  min-height: 100px;
}

.grid-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1.5rem;
}

.btn-submit,
.btn-switch {
  width: 100%;
  padding: 0.75rem;
  border: none;
  border-radius: 0.5rem;
  font-size: 1rem;
  cursor: pointer;
  transition: background 0.3s;
  margin-bottom: 1.5rem;
}

.btn-submit {
  background: #007bff;
  color: white;
}

.btn-submit:hover {
  background: #0069d9;
}

.btn-switch {
  background: transparent;
  margin-top: 1rem;
  color: #007bff;
}

.btn-switch:hover {
  color: #0056b3;
}

.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s;
}

.fade-enter,
.fade-leave-to {
  opacity: 0;
}
</style>
