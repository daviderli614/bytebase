<template>
  <div
    class="flex-1 flex flex-col gap-2 relative overflow-auto"
    style="width: calc(100vw - 8rem); max-width: 50rem"
  >
    <div
      class="flex flex-col items-start gap-2 sm:flex-row sm:justify-between sm:items-center"
    >
      <div class="flex justify-start items-center gap-2">
        <NButton @click="state.transfer.show = true">
          <template #icon>
            <ChevronsDownIcon class="w-4 h-4" />
          </template>
          {{ $t("quick-action.transfer-in-db") }}
        </NButton>
        <NButton v-if="showSettingButton" @click="state.setting.show = true">
          <template #icon>
            <SettingsIcon class="w-4 h-4" />
          </template>
          {{ $t("common.setting") }}
        </NButton>
        <NButton v-if="showMembersButton" @click="state.members.show = true">
          <template #icon>
            <UsersIcon class="w-4 h-4" />
          </template>
          {{ $t("common.members") }}
        </NButton>
      </div>
    </div>

    <div class="relative">
      <DatabaseV1Table
        v-if="ready"
        mode="PROJECT_SHORT"
        :database-list="databaseList"
        :show-selection="false"
        :custom-click="true"
        @row-click="handleDatabaseClick"
      />
    </div>
    <MaskSpinner v-if="!ready" />

    <Drawer v-model:show="state.transfer.show">
      <TransferDatabaseForm
        :project-name="project.name"
        :on-success="() => (state.transfer.show = false)"
        @dismiss="state.transfer.show = false"
      />
    </Drawer>

    <Drawer v-model:show="state.setting.show">
      <DrawerContent
        :title="$t('common.setting')"
        body-style="width: 40vw; max-width: 600px; min-width: 320px;"
      >
        <ProjectSettingPanel
          :project="project"
          :allow-edit="allowEditSetting"
        />
      </DrawerContent>
    </Drawer>

    <Drawer v-model:show="state.members.show">
      <DrawerContent
        :title="$t('common.members')"
        body-style="width: 60vw; max-width: calc(100vw - 4rem); min-width: 400px;"
      >
        <ProjectMemberPanel :project="project" :allow-edit="allowEditMembers" />
      </DrawerContent>
    </Drawer>
  </div>
</template>

<script setup lang="ts">
import { ChevronsDownIcon, SettingsIcon, UsersIcon } from "lucide-vue-next";
import { NButton } from "naive-ui";
import { computed, reactive } from "vue";
import { useRouter } from "vue-router";
import ProjectMemberPanel from "@/components/ProjectMember/ProjectMemberPanel.vue";
import ProjectSettingPanel from "@/components/ProjectSettingPanel.vue";
import { TransferDatabaseForm } from "@/components/TransferDatabaseForm";
import MaskSpinner from "@/components/misc/MaskSpinner.vue";
import { Drawer, DrawerContent } from "@/components/v2";
import DatabaseV1Table from "@/components/v2/Model/DatabaseV1Table";
import { useDatabaseV1List } from "@/store/modules/v1/databaseList";
import {
  DEFAULT_PROJECT_NAME,
  type ComposedDatabase,
  type ComposedProject,
} from "@/types";
import { State } from "@/types/proto/v1/common";
import { autoDatabaseRoute, hasProjectPermissionV2 } from "@/utils";

type LocalState = {
  transfer: {
    show: boolean;
  };
  setting: {
    show: boolean;
  };
  members: {
    show: boolean;
  };
};

const props = defineProps<{
  project: ComposedProject;
}>();

const router = useRouter();
const { databaseList, ready } = useDatabaseV1List(props.project.name);
const state = reactive<LocalState>({
  transfer: { show: false },
  setting: { show: false },
  members: { show: false },
});

const showSettingButton = computed(() => {
  if (props.project.name === DEFAULT_PROJECT_NAME) return false;
  return hasProjectPermissionV2(props.project, "bb.projects.get");
});
const allowEditSetting = computed(() => {
  if (props.project.name === DEFAULT_PROJECT_NAME) return false;
  if (props.project.state === State.DELETED) {
    return false;
  }
  return hasProjectPermissionV2(props.project, "bb.projects.update");
});

const showMembersButton = computed(() => {
  if (props.project.name === DEFAULT_PROJECT_NAME) return false;
  return (
    hasProjectPermissionV2(props.project, "bb.projects.get") &&
    hasProjectPermissionV2(props.project, "bb.projects.getIamPolicy")
  );
});
const allowEditMembers = computed(() => {
  if (props.project.name === DEFAULT_PROJECT_NAME) return false;
  if (props.project.state === State.DELETED) {
    return false;
  }

  return hasProjectPermissionV2(props.project, "bb.projects.setIamPolicy");
});

const handleDatabaseClick = (event: MouseEvent, database: ComposedDatabase) => {
  router.push(autoDatabaseRoute(router, database));
};
</script>
