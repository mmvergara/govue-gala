<template>
  <main class="pt-[5vh]">
    <h1 class="text-lg font-semibold text-center pb-6">Create Post</h1>
    <form
      v-on:submit="submitPost"
      class="flex flex-col justify-center items-center"
    >
      <div
        class="flex flex-col justify-center items-center gap-1 border-2 p-6 border-retroPurple rounded-sm w-full max-w-[500px]"
      >
        <label for="postTitle" class="w-full">Post Title</label>
        <input
          v-model="postTitle"
          type="text"
          id="postTitle"
          class="bg-retroCream p-2 rounded-sm w-full mb-4"
        />
        <label for="postDescription" class="w-full">Post Description</label>
        <input
          v-model="postDescription"
          type="text"
          id="postDescription"
          class="bg-retroCream p-4 rounded-sm w-full mb-6"
        />
        <button
          type="submit"
          class="bg-retroPurplePink p-4 rounded-sm text-white flex-auto w-full"
        >
          Create Post
        </button>
      </div>
    </form>
  </main>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { postRequest } from "~/lib/fetching";
const router = useRouter();
const postTitle = ref("");
const postDescription = ref("");

const submitPost = async (e: Event) => {
  e.preventDefault();
  const { error } = await postRequest("/post", {
    author_id: "43637313-4e92-4666-8174-366226658cfe",
    post_title: postTitle.value,
    post_description: postDescription.value,
  });
  if (error) {
    useNuxtApp().$toast.error("Post created successfully!");
    return;
  }
  useNuxtApp().$toast.success("Post created successfully!");
  setTimeout(() => {
    router.push("/");
  }, 2000);
};
</script>
