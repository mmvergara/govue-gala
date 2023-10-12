<template>
  <main class="pt-[5vh]">
    <h1 v-if="isSignup" class="text-lg font-semibold text-center pb-6">
      Sign up | Golang Gala
    </h1>
    <h1 v-else class="text-lg font-semibold text-center pb-6">
      Sign in | Golang Gala
    </h1>
    <form
      v-on:submit="authSubmit"
      class="flex flex-col justify-center items-center"
    >
      <div
        class="flex flex-col justify-center items-center gap-1 border-t-2 border-b-2 p-6 border-retroPurple rounded-sm w-full max-w-[500px]"
      >
        <label for="authUsername" class="w-full">Username</label>
        <input
          v-model="username"
          type="text"
          id="authUsername"
          class="bg-retroCream p-2 rounded-sm w-full"
        />
        <label for="authPassword" class="w-full">Password</label>
        <input
          v-model="password"
          type="password"
          id="authPassword"
          class="bg-retroCream p-2 rounded-sm w-full mb-4"
        />
        <button
          v-if="isSignup"
          type="submit"
          class="bg-retroPurplePink p-4 rounded-sm text-white flex-auto w-full"
        >
          Sign up
        </button>
        <button
          v-else
          type="submit"
          class="bg-retroPurplePink p-4 rounded-sm text-white flex-auto w-full"
        >
          Sign in
        </button>
      </div>
    </form>
    <button
      v-if="isSignup"
      @click="isSignup = !isSignup"
      class="text-center w-full pt-4 underline underline-offset-4"
    >
      I already have an account
    </button>
    <button
      v-else
      @click="isSignup = !isSignup"
      class="text-center w-full pt-4 underline underline-offset-4"
    >
      Create New Account
    </button>
  </main>
</template>

<script setup lang="ts">
import { ref } from "vue";

const username = ref("");
const password = ref("");

const isSignup = ref(false);

const authSubmit = (e: Event) => {
  e.preventDefault();
  if (isSignup.value) {
    console.log("signup");
  } else {
    console.log("signin");
  }
};
</script>
