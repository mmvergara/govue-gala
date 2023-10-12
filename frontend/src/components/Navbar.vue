<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
  <nav class="flex bg-retroPurple gap-2 text-white p-2">
    <RouterLink to="/" class="mr-auto nav-link">Govue Gala</RouterLink>
    <RouterLink to="/create/post" class="nav-link">Create Post</RouterLink>
  </nav>
</template>

<style scoped>
@tailwind components;

@layer components {
  .nav-link {
    @apply hover:bg-gray-400 hover:bg-opacity-20 p-2 rounded-sm;
  }
}
</style>
