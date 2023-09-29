<template>
  <nav class="flex bg-retroPurple gap-2 text-white p-2">
    <nuxt-link to="/" class="mr-auto nav-link">Golang Gala</nuxt-link>
    <nuxt-link to="/post/create" class="nav-link">Create Post</nuxt-link>
    <nuxt-link to="/auth" class="nav-link">Sign In</nuxt-link>
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
