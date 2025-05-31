<!-- Developed By Taipei Urban Intelligence Center 2023-2024 -->
<!-- 
Lead Developer:  Igor Ho (Full Stack Engineer)
Data Pipelines:  Iima Yu (Data Scientist)
Design and UX: Roy Lin (Fmr. Consultant), Chu Chen (Researcher)
Systems: Ann Shih (Systems Engineer)
Testing: Jack Huang (Data Scientist), Ian Huang (Data Analysis Intern) 
-->
<!-- Department of Information Technology, Taipei City Government -->

<script setup>
import DashboardComponent from "../dashboardComponent/DashboardComponent.vue";
import router from "../router";
import { useContentStore } from "../store/contentStore";
import { useDialogStore } from "../store/dialogStore";
import { useAuthStore } from "../store/authStore";

import MoreInfo from "../components/dialogs/MoreInfo.vue";
import ReportIssue from "../components/dialogs/ReportIssue.vue";
import ImageTextView from "./ImageTextView.vue";
import { ref, onMounted, onUnmounted } from "vue";

const contentStore = useContentStore();
const dialogStore = useDialogStore();
const authStore = useAuthStore();

const showImageText = ref(true); // 控制是否顯示右側圖片文字頁
const leftPaneWidth = ref(75); // 百分比，預設 60%
let isDragging = false;
let dragStartX = 0;
let dragStartWidth = 0;

function handleOpenSettings() {
	contentStore.editDashboard = JSON.parse(
		JSON.stringify(contentStore.currentDashboard)
	);
	dialogStore.addEdit = "edit";
	dialogStore.showDialog("addEditDashboards");
}

function toggleFavorite(id) {
	if (contentStore.favorites.components.includes(id)) {
		contentStore.unfavoriteComponent(id);
	} else {
		contentStore.favoriteComponent(id);
	}
}
function handleMoreInfo(item) {
	if (authStore.isMobileDevice && authStore.isNarrowDevice) {
		router.push({
			name: "component-info",
			params: { index: item.index },
		});
	} else {
		dialogStore.showMoreInfo(item);
	}
}

function onSplitterMouseDown(e) {
	isDragging = true;
	dragStartX = e.touches ? e.touches[0].clientX : e.clientX;
	dragStartWidth = leftPaneWidth.value;
	document.body.style.cursor = "col-resize";
}

function onSplitterMouseMove(e) {
	if (!isDragging) return;
	const split = document.querySelector(".dashboard-split");
	if (!split) return;
	const rect = split.getBoundingClientRect();
	let clientX = e.touches ? e.touches[0].clientX : e.clientX;
	// 計算滑鼠移動的距離
	let delta = clientX - dragStartX;
	let percentDelta = (delta / rect.width) * 100;
	let percent = dragStartWidth + percentDelta;
	percent = Math.max(50, Math.min(75, percent)); // 限制最小/最大寬度
	leftPaneWidth.value = percent;
}

function onSplitterMouseUp() {
	isDragging = false;
	document.body.style.cursor = "";
}

onMounted(() => {
	window.addEventListener("mousemove", onSplitterMouseMove);
	window.addEventListener("mouseup", onSplitterMouseUp);
	window.addEventListener("touchmove", onSplitterMouseMove);
	window.addEventListener("touchend", onSplitterMouseUp);
});
onUnmounted(() => {
	window.removeEventListener("mousemove", onSplitterMouseMove);
	window.removeEventListener("mouseup", onSplitterMouseUp);
	window.removeEventListener("touchmove", onSplitterMouseMove);
	window.removeEventListener("touchend", onSplitterMouseUp);
});
</script>

<template>
	<div v-if="showImageText" class="dashboard-split">
		<div
			class="dashboard-split-left"
			:style="{
				flexBasis: leftPaneWidth + '%',
				width: leftPaneWidth + '%',
				minWidth: '0',
			}"
		>
			<!-- 1. If the dashboard is map-layers -->
			<div
				v-if="
					contentStore.currentDashboard.index?.includes('map-layers')
				"
				class="dashboard"
			>
				<DashboardComponent
					v-for="item in contentStore.currentDashboard.components"
					:key="`${item.index}-${item.city}`"
					:config="item"
					mode="half"
					:info-btn="true"
					:active-city="item.city"
					:select-btn="true"
					:select-btn-disabled="
						contentStore.cityManager.getSelectList(
							contentStore.currentDashboard?.city
						).length === 1
					"
					:select-btn-list="
						contentStore.cityManager.getSelectList(
							contentStore.currentDashboard?.city
						)
					"
					:city-tag="
						contentStore.cityManager.getTagList(
							contentStore.currentDashboard?.city
						)
					"
					:favorite-btn="authStore.token ? true : false"
					:is-favorite="
						contentStore.favorites?.components.includes(item.id)
					"
					@favorite="
						(id) => {
							toggleFavorite(id);
						}
					"
					@info="
						(item) => {
							handleMoreInfo(item);
						}
					"
					@change-city="
						(city) => {
							const selectedData =
								contentStore.cityDashboard.components.find(
									(data) => {
										if (
											data.index === item.index &&
											data.city === city
										) {
											return data;
										}
									}
								);
							const componentIndex =
								contentStore.currentDashboard.components.findIndex(
									(item) => item.id === selectedData.id
								);
							if (selectedData) {
								contentStore.setComponentData(
									componentIndex,
									selectedData
								);
							}
						}
					"
				/>
				<MoreInfo />
				<ReportIssue />
			</div>
			<!-- 2. Dashboards that have components -->
			<div
				v-else-if="
					contentStore.currentDashboard.components?.length !== 0 ||
					contentStore.cityDashboard.components?.length !== 0
				"
				class="dashboard"
			>
				<DashboardComponent
					v-for="item in contentStore.currentDashboard.components"
					:key="`${item.index}-${item.city}`"
					:config="item"
					:info-btn="true"
					:active-city="item.city"
					:select-btn="true"
					:select-btn-disabled="
						contentStore.cityManager.getSelectList(
							contentStore.currentDashboard?.city
						).length === 1 ||
						contentStore.currentDashboardExcluded.components.filter(
							(data) => data.index === item.index
						).length === 0
					"
					:select-btn-list="
						contentStore.currentDashboard?.city
							? contentStore.cityManager.getSelectList(
									contentStore.currentDashboard?.city
							  )
							: contentStore.cityManager.getCities(
									contentStore.cityManager.activeCities
							  )
					"
					:city-tag="
						contentStore.currentDashboard?.city
							? contentStore.cityManager.getTagList(
									contentStore.currentDashboard?.city
							  )
							: contentStore.cityManager.getTagList(item.city)
					"
					:delete-btn="
						contentStore.personalDashboards
							.map((item) => item.index)
							.includes(contentStore.currentDashboard.index)
					"
					:favorite-btn="
						authStore.token &&
						contentStore.currentDashboard.icon !== 'favorite'
					"
					:is-favorite="
						contentStore.favorites?.components.includes(item.id)
					"
					@favorite="
						(id) => {
							toggleFavorite(id);
						}
					"
					@info="
						(item) => {
							handleMoreInfo(item);
						}
					"
					@delete="
						(id) => {
							contentStore.deleteComponent(id);
						}
					"
					@change-city="
						(city) => {
							const selectedData =
								contentStore.cityDashboard.components.find(
									(data) => {
										if (
											data.index === item.index &&
											data.city === city
										) {
											return data;
										}
									}
								);
							const componentIndex =
								contentStore.currentDashboard.components.findIndex(
									(item) => item.id === selectedData.id
								);
							if (selectedData) {
								contentStore.setComponentData(
									componentIndex,
									selectedData
								);
							}
						}
					"
				/>
				<MoreInfo />
				<ReportIssue />
			</div>
			<!-- 3. If dashboard is still loading -->
			<div
				v-else-if="contentStore.loading"
				class="dashboard dashboard-nodashboard"
			>
				<div class="dashboard-nodashboard-content">
					<div />
				</div>
			</div>
			<!-- 4. If dashboard failed to load -->
			<div
				v-else-if="contentStore.error"
				class="dashboard dashboard-nodashboard"
			>
				<div class="dashboard-nodashboard-content">
					<span>sentiment_very_dissatisfied</span>
					<h2>發生錯誤，無法載入儀表板</h2>
				</div>
			</div>
			<!-- 5. Dashboards that don't have components -->
			<div v-else class="dashboard dashboard-nodashboard">
				<div class="dashboard-nodashboard-content">
					<span>addchart</span>
					<h2>尚未加入組件</h2>
					<button
						v-if="contentStore.currentDashboard.icon !== 'favorite'"
						class="hide-if-mobile"
						@click="handleOpenSettings"
					>
						加入您的第一個組件
					</button>
					<p v-else>點擊其他儀表板組件之愛心以新增至收藏組件</p>
				</div>
			</div>
		</div>
		<div
			class="splitter"
			@mousedown="onSplitterMouseDown"
			@touchstart.prevent="onSplitterMouseDown"
		>
			<!-- Center icon button -->
			<button
				class="splitter-toggle-btn"
				@click="showImageText = false"
				:title="'收合右側'"
			>
				<span>chevron_right</span>
			</button>
		</div>
		<div
			class="dashboard-split-right"
			:style="{
				flexBasis: 100 - leftPaneWidth + '%',
				width: 100 - leftPaneWidth + '%',
				minWidth: '0',
			}"
		>
			<ImageTextView />
		</div>
	</div>
	<template v-else>
		<!-- 1. If the dashboard is map-layers -->
		<div
			v-if="contentStore.currentDashboard.index?.includes('map-layers')"
			class="dashboard"
		>
			<DashboardComponent
				v-for="item in contentStore.currentDashboard.components"
				:key="`${item.index}-${item.city}`"
				:config="item"
				mode="half"
				:info-btn="true"
				:active-city="item.city"
				:select-btn="true"
				:select-btn-disabled="
					contentStore.cityManager.getSelectList(
						contentStore.currentDashboard?.city
					).length === 1
				"
				:select-btn-list="
					contentStore.cityManager.getSelectList(
						contentStore.currentDashboard?.city
					)
				"
				:city-tag="
					contentStore.cityManager.getTagList(
						contentStore.currentDashboard?.city
					)
				"
				:favorite-btn="authStore.token ? true : false"
				:is-favorite="
					contentStore.favorites?.components.includes(item.id)
				"
				@favorite="
					(id) => {
						toggleFavorite(id);
					}
				"
				@info="
					(item) => {
						handleMoreInfo(item);
					}
				"
				@change-city="
					(city) => {
						const selectedData =
							contentStore.cityDashboard.components.find(
								(data) => {
									if (
										data.index === item.index &&
										data.city === city
									) {
										return data;
									}
								}
							);

						const componentIndex =
							contentStore.currentDashboard.components.findIndex(
								(item) => item.id === selectedData.id
							);

						if (selectedData) {
							contentStore.setComponentData(
								componentIndex,
								selectedData
							);
						}
					}
				"
			/>
			<MoreInfo />
			<ReportIssue />
		</div>
		<!-- 2. Dashboards that have components -->
		<div
			v-else-if="
				contentStore.currentDashboard.components?.length !== 0 ||
				contentStore.cityDashboard.components?.length !== 0
			"
			class="dashboard"
		>
			<DashboardComponent
				v-for="item in contentStore.currentDashboard.components"
				:key="`${item.index}-${item.city}`"
				:config="item"
				:info-btn="true"
				:active-city="item.city"
				:select-btn="true"
				:select-btn-disabled="
					contentStore.cityManager.getSelectList(
						contentStore.currentDashboard?.city
					).length === 1 ||
					contentStore.currentDashboardExcluded.components.filter(
						(data) => data.index === item.index
					).length === 0
				"
				:select-btn-list="
					contentStore.currentDashboard?.city
						? contentStore.cityManager.getSelectList(
								contentStore.currentDashboard?.city
						  )
						: contentStore.cityManager.getCities(
								contentStore.cityManager.activeCities
						  )
				"
				:city-tag="
					contentStore.currentDashboard?.city
						? contentStore.cityManager.getTagList(
								contentStore.currentDashboard?.city
						  )
						: contentStore.cityManager.getTagList(item.city)
				"
				:delete-btn="
					contentStore.personalDashboards
						.map((item) => item.index)
						.includes(contentStore.currentDashboard.index)
				"
				:favorite-btn="
					authStore.token &&
					contentStore.currentDashboard.icon !== 'favorite'
				"
				:is-favorite="
					contentStore.favorites?.components.includes(item.id)
				"
				@favorite="
					(id) => {
						toggleFavorite(id);
					}
				"
				@info="
					(item) => {
						handleMoreInfo(item);
					}
				"
				@delete="
					(id) => {
						contentStore.deleteComponent(id);
					}
				"
				@change-city="
					(city) => {
						const selectedData =
							contentStore.cityDashboard.components.find(
								(data) => {
									if (
										data.index === item.index &&
										data.city === city
									) {
										return data;
									}
								}
							);

						const componentIndex =
							contentStore.currentDashboard.components.findIndex(
								(item) => item.id === selectedData.id
							);

						if (selectedData) {
							contentStore.setComponentData(
								componentIndex,
								selectedData
							);
						}
					}
				"
			/>
			<MoreInfo />
			<ReportIssue />
			<div class="float-btn">
				<button
					class="splitter-toggle-btn"
					@click="showImageText = true"
					:title="'展開右側'"
				>
					<span>chevron_left</span>
				</button>
			</div>
		</div>
		<!-- 3. If dashboard is still loading -->
		<div
			v-else-if="contentStore.loading"
			class="dashboard dashboard-nodashboard"
		>
			<div class="dashboard-nodashboard-content">
				<div />
			</div>
		</div>
		<!-- 4. If dashboard failed to load -->
		<div
			v-else-if="contentStore.error"
			class="dashboard dashboard-nodashboard"
		>
			<div class="dashboard-nodashboard-content">
				<span>sentiment_very_dissatisfied</span>
				<h2>發生錯誤，無法載入儀表板</h2>
			</div>
		</div>
		<!-- 5. Dashboards that don't have components -->
		<div v-else class="dashboard dashboard-nodashboard">
			<div class="dashboard-nodashboard-content">
				<span>addchart</span>
				<h2>尚未加入組件</h2>
				<button
					v-if="contentStore.currentDashboard.icon !== 'favorite'"
					class="hide-if-mobile"
					@click="handleOpenSettings"
				>
					加入您的第一個組件
				</button>
				<p v-else>點擊其他儀表板組件之愛心以新增至收藏組件</p>
			</div>
		</div>
	</template>
</template>

<style scoped lang="scss">
.dashboard {
	max-height: calc(100vh - 127px);
	max-height: calc(var(--vh) * 100 - 127px);
	display: grid;
	row-gap: var(--font-s);
	column-gap: var(--font-s);
	margin: var(--font-m) var(--font-m);
	overflow-y: scroll;

	@media (min-width: 720px) {
		grid-template-columns: 1fr 1fr;
	}

	@media (min-width: 1200px) {
		grid-template-columns: 1fr 1fr 1fr;
	}

	@media (min-width: 1800px) {
		grid-template-columns: 1fr 1fr 1fr 1fr;
	}

	@media (min-width: 2200px) {
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
	}

	&-nodashboard {
		grid-template-columns: 1fr;

		&-content {
			width: 100%;
			height: calc(100vh - 127px);
			height: calc(var(--vh) * 100 - 127px);
			display: flex;
			flex-direction: column;
			align-items: center;
			justify-content: center;

			span {
				margin-bottom: var(--font-ms);
				font-family: var(--font-icon);
				font-size: 2rem;
			}

			button {
				color: var(--color-highlight);
			}

			div {
				width: 2rem;
				height: 2rem;
				border-radius: 50%;
				border: solid 4px var(--color-border);
				border-top: solid 4px var(--color-highlight);
				animation: spin 0.7s ease-in-out infinite;
			}
		}
	}
}
.dashboard-split {
	display: flex;
	height: 100%;
	position: relative;

	span {
		margin-bottom: var(--font-ms);
		font-family: var(--font-icon);
		font-size: 2rem;
	}
	.dashboard-split-left {
		flex: 1 1 0;
		min-width: 0;
		border-right: 1px solid var(--color-bg);
		background: var(--color-bg);
		overflow-y: auto;
	}
	.splitter {
		width: 8px;
		cursor: col-resize;
		background-color: var(--color-border);
		z-index: 2;
		transition: background 0.2s;
		user-select: none;
		position: relative;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.splitter:hover,
	.splitter:active {
		background: var(--color-highlight);
	}
	.dashboard-split-right {
		padding-inline: 0.5rem;
		flex: 1 1 0;
		min-width: 0;
		background: var(--color-bg);
		overflow-y: auto;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
}
.float-btn {
	position: absolute;
	top: 50%;
	right: 0.5rem;
	height: 2rem;
	width: 2rem;
	z-index: 3;
}
.splitter-toggle-btn {
	position: fixed;
	height: 2rem;
	width: 2rem;
	background: #fff;
	border-radius: 50%;
	// box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
	// padding: 2px 4px;
	cursor: pointer;
	z-index: 5;

	span {
		font-size: 2rem;
		font-family: var(--font-icon);
		color: var(--color-highlight);
	}
}

@keyframes spin {
	to {
		transform: rotate(360deg);
	}
}
</style>
