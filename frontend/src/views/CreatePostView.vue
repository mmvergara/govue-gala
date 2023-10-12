<template>
  <main class="pt-[5vh]">
    <h1 class="text-lg font-semibold text-center pb-6">Create Post</h1>
    <form
      v-on:submit="submitPost"
      class="flex flex-col justify-center items-center"
    >
      <div
        class="flex flex-col justify-center items-center gap-1 border-t-2 border-b-2 p-6 border-retroPurple rounded-sm w-full max-w-[500px]"
      >
        <label for="postTitle" class="w-full">Post Title</label>
        <input
          v-model="postTitle"
          type="text"
          id="postTitle"
          class="bg-retroCream p-2 rounded-sm w-full mb-4"
        />
        <label for="postDescription" class="w-full">Post Description</label>
        <textarea
          v-model="postDescription"
          type="text"
          id="postDescription"
          class="bg-retroCream p-4 rounded-sm w-full mb-6"
        ></textarea>
        <button
          type="submit"
          class="bg-retroPurplePink p-4 text-white flex-auto w-full rounded-md"
        >
          Create Post
        </button>
        <p class="text-red-500">{{ err }}</p>
      </div>
    </form>
  </main>
</template>

<script setup lang="ts">
import { AxiosRequest } from "@/components/AxiosInstance";
import { ref } from "vue";
const postTitle = ref("");
const postDescription = ref("");
const err = ref("");

const submitPost = async (e: Event) => {
  e.preventDefault();

  const { error } = await AxiosRequest("/post", "POST", {
    author_id: "fa025492-9e1b-4d7e-8af6-5a7b2a16d359",
    post_title: postTitle.value,
    post_description: postDescription.value,
  });

  if (error) {
    err.value = error;
  } else {
    err.value = "";
    postTitle.value = "";
    postDescription.value = "";
  }
};
</script>
