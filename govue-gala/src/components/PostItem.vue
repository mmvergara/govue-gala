<template>
  <section>
    <h2>{{ post.post_title }}</h2>
    <p>{{ post.post_description }}</p>
  </section>
</template>

<script setup lang="ts">
import { defineProps } from "vue";
import type { PostType } from "../types/post-types";
const { post } = defineProps<{ post: PostType }>();
</script>
