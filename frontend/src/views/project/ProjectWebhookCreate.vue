<template>
  <ProjectWebhookForm
    :create="true"
    :project="project"
    :webhook="defaultNewWebhook"
  />
</template>

<script lang="ts" setup>
import { computed } from "vue";
import ProjectWebhookForm from "@/components/ProjectWebhookForm.vue";
import { useProjectByName } from "@/store";
import { projectNamePrefix } from "@/store/modules/v1/common";
import { emptyProjectWebhook } from "@/types";

const props = defineProps<{
  projectId: string;
}>();

const { project } = useProjectByName(
  computed(() => `${projectNamePrefix}${props.projectId}`)
);

const defaultNewWebhook = emptyProjectWebhook();
</script>
