<template>
	<div class="floating-options-toggle">
		<div class="floating-dropdown" v-if="showOptions">
			<select v-model="selectedCity">
				<option value="臺北">臺北</option>
				<option value="新北">新北</option>
			</select>
			<select v-model="selectedDistrict">
				<option
					v-for="district in districts"
					:key="district"
					:value="district"
				>
					{{ district }}
				</option>
			</select>
			<select v-model="selectedMarket">
				<option v-for="market in markets" :key="market" :value="market">
					{{ market }}
				</option>
			</select>
		</div>
		<button class="floating-btn" @click="showOptions = !showOptions">
			<span v-if="!showOptions">arrow_drop_down</span>
			<span v-else>arrow_drop_up</span>
		</button>
	</div>
	<DialogContainer :dialog="`addAds`" @on-close="handleClose">
		<div class="dialog-content">
			<h2>新增活動</h2>
			<label for="imageUrl">圖片網址:</label>
			<input type="text" id="imageUrl" v-model="imageUrl" />
			<label for="linkUrl">連結網址:</label>
			<input type="text" id="linkUrl" v-model="linkUrl" />
			<div class="dialog-content-control">
				<button
					v-if="imageUrl && linkUrl"
					class="dialog-content-control-confirm"
					@click="addAds"
				>
					提交
				</button>
			</div>
		</div>
	</DialogContainer>
	<div class="image-text-view">
		<div class="image-section scrollable-images">
			<div
				v-if="
					marketImages[selectedMarket] &&
					marketImages[selectedMarket].length > 0
				"
				:style="{ overflowY: 'auto' }"
			>
				<div class="title">
					<h1>{{ selectedMarket }}{{ title }}</h1>
				</div>
				<div
					v-for="(img, idx) in marketImages[selectedMarket]"
					:key="img.url"
					class="image-row"
				>
					<a v-if="img.link" :href="img.link" target="_blank"
						><img :src="img.url" :alt="selectedMarket + ' 圖片'"
					/></a>
					<img
						v-else
						:src="img.url"
						:alt="selectedMarket + ' 圖片'"
					/>
					<div
						v-if="authStore.editAds"
						class="delete-btn"
						@click="deleteAd(idx)"
					>
						<span>delete_forever</span>
					</div>
				</div>
				<div
					v-if="authStore.editAds"
					class="add-btn"
					@click="dialogStore.showDialog('addAds')"
				>
					<span>add</span>
				</div>
			</div>
			<div v-else>
				<img
					:src="imageDefaultUrl"
					alt="展示圖片"
					:style="{ maxWidth: '100%' }"
				/>
			</div>
		</div>
	</div>
</template>

<script setup>
import { ref, watch } from "vue";
import { useAuthStore } from "../store/authStore";
import { useDialogStore } from "../store/dialogStore";

import DialogContainer from "../components/dialogs/DialogContainer.vue";
const imageDefaultUrl = ref("https://placehold.co/600x400?text=展示圖片");
const title = ref("最新活動");

const authStore = useAuthStore();
const dialogStore = useDialogStore();

const imageUrl = ref("");
const linkUrl = ref("");

const selectedCity = ref("臺北");
const selectedDistrict = ref("");
const selectedMarket = ref("");
const showOptions = ref(false);

function handleClose() {
	dialogStore.hideAllDialogs();
}
function addAds() {
	if (!imageUrl.value) {
		dialogStore.showNotification("error", "請輸入圖片網址");
		return;
	}
	if (!marketImages.value[selectedMarket.value]) {
		marketImages.value[selectedMarket.value] = [];
	}
	marketImages.value[selectedMarket.value].push({
		url: imageUrl.value,
		link: linkUrl.value,
	});
	imageUrl.value = "";
	linkUrl.value = "";
	dialogStore.showNotification("success", "活動新增成功");
	dialogStore.hideAllDialogs();
}
const districtsMap = {
	臺北: [
		"中正區",
		"中山區",
		"信義區",
		"內湖區",
		"北投區",
		"士林區",
		"大同區",
		"大安區",
		"文山區",
		"松山區",
		"萬華區",
	],
	新北: [
		"板橋區",
		"深坑區",
		"鶯歌區",
		"蘆洲區",
		"石碇區",
		"永和區",
		"平溪區",
		"淡水區",
		"新莊區",
		"三峽區",
		"瑞芳區",
		"金山區",
		"雙溪區",
		"萬里區",
		"烏來區",
		"坪林區",
		"八里區",
		"三重區",
		"林口區",
	],
};
const marketsMap = {
	// 臺北
	中山區: [
		"中山北路",
		"南西心中山",
		"四平陽光(商圈)",
		"大正町商圈",
		"大直(商圈)",
		"晴光(商圈)",
		"條通(商圈)",
		"民族濱江汽車",
	],
	中正區: [
		"臺大公館(商圈)",
		"中華路影音(商圈)",
		"北門相機(商圈)",
		"南昌家具(商圈)",
		"大光華(商圈)",
		"愛國東路婚紗(商圈)",
		"榮町商圈",
		"沅陵街(商圈)",
		"重慶南路書店(商圈)",
	],
	信義區: ["五分埔(商圈)", "吳興街(商圈)"],
	內湖區: ["內湖737(商圈)", "西湖(商圈)"],
	北投區: ["新北投溫泉(商圈)", "石牌捷運(商圈)", "行義路溫泉美食(商圈)"],
	士林區: [
		"士林捷運商圈",
		"士林觀光夜市(商圈)",
		"天母(商圈)",
		"承德路中古汽車(商圈)",
		"蘭雅(商圈)",
	],
	大同區: [
		"圓山(商圈)",
		"圓環太平",
		"大龍峒(商圈)",
		"寧夏夜市(商圈)",
		"後站(商圈)",
		"朝陽服飾材料(商圈)",
		"臺北大橋頭延三(商圈)",
		"華陰街(商圈)",
		"赤峰",
		"迪化街(商圈)",
	],
	大安區: ["文昌家具(商圈)", "東區商圈", "永康(商圈)", "龍泉(商圈)"],
	文山區: ["萬芳(商圈)", "貓空(商圈)"],
	松山區: ["民生社區"],
	萬華區: ["加蚋(商圈)", "艋舺(商圈)", "艋舺(夜市)", "萬華", "西門町(商圈)"],
	// 新北
	深坑區: ["新北市商圈聯合發展協會", "深坑老街商圈"],
	鶯歌區: ["鶯歌商圈"],
	蘆洲區: ["蘆洲廟口商圈", "猴硐商圈"],
	石碇區: ["石碇商圈"],
	永和區: ["韓國街商圈"],
	平溪區: ["平溪魅力商圈", "菁桐商圈"],
	淡水區: ["淡水老街商圈"],
	板橋區: ["板橋府中商圈", "亞東商圈"],
	新莊區: ["新莊廟街商圈", "新莊副都心商圈"],
	三峽區: ["三峽民權老街商圈", "三峽三角湧商圈", "三峽秀川商圈"],
	瑞芳區: ["九份老街商圈", "瑞芳老街商圈"],
	金山區: ["金山商圈"],
	雙溪區: ["雙溪商圈"],
	萬里區: ["萬里商圈", "野柳商圈"],
	烏來區: ["烏來商圈"],
	坪林區: ["坪林商圈"],
	八里區: ["八里左岸商圈", "八里渡船頭商圈"],
	三重區: ["三重碧華布街商圈"],
	林口區: ["林口建城商圈"],
};

// 商圈圖片對應表
const marketImages = ref({
	// 臺大公館(商圈)
	"臺大公館(商圈)": [
		{
			url: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic.accupass.com%2Feventbanner%2F2109290321521979163684.jpg&f=1&nofb=1&ipt=43adabc270d2aa2ad1519c6932a194761415a8b44edc0a51593bd887ec7ac49e",
			link: "https://www.accupass.com/event/2109290326457002327540",
		},
		{
			url: "https://scontent.ftpe7-2.fna.fbcdn.net/v/t39.30808-6/458084404_1051127430350160_6110758109483297383_n.png?_nc_cat=109&ccb=1-7&_nc_sid=cc71e4&_nc_ohc=EoVRd0q6m1IQ7kNvwE89lEE&_nc_oc=AdnhYEn5HATfNVg9qXAbn_IV2Bf9Fcg3U6d5EdEAz7gWf5hTXrudY-WZADDwfdqMgno&_nc_zt=23&_nc_ht=scontent.ftpe7-2.fna&_nc_gid=13JpKd7frKXZJ2FMPY4Sgw&oh=00_AfLwAtAFpAacKEUxmapX95EZKP8nZvGIYntieC3yawvWtg&oe=6840D523",
			link: "",
		},
		{
			url: "https://scontent.ftpe7-1.fna.fbcdn.net/v/t39.30808-6/494021483_1252416143554620_1016594463807232717_n.jpg?_nc_cat=106&ccb=1-7&_nc_sid=127cfc&_nc_ohc=pApWQ88wDH8Q7kNvwFkJAp3&_nc_oc=AdlILU4XPGM9nHmAnW4u23K5KBTs5J04KclJuVGRJZ1o1LyCjO3O1I8bwpqPOgkGx3c&_nc_zt=23&_nc_ht=scontent.ftpe7-1.fna&_nc_gid=_Xn9s58Nr5ymuwUXm7244g&oh=00_AfI286bTfZ5Gw3QY2-QhlEI_Mq6Qg7OqKelfGHK-J4Q-Lg&oe=6840B464",
			link: "",
		},
		{
			url: "https://scontent.ftpe7-1.fna.fbcdn.net/v/t39.30808-6/492211952_1251790406950527_943383893772262145_n.jpg?_nc_cat=106&ccb=1-7&_nc_sid=127cfc&_nc_ohc=okns-ylD_SsQ7kNvwFDEZmH&_nc_oc=AdnKaiQsfck5thdU0KIstns4W7lOHp-57IGa-QsSlJg0w-Q2zb6O7UcIkIfljvnbAaM&_nc_zt=23&_nc_ht=scontent.ftpe7-1.fna&_nc_gid=KA9-Zwjl6_A4dxQjpmoehQ&oh=00_AfKtizZsbtwg7YkltvKdRx_zaxrVp8y4ejP8Y_QEX-eJnA&oe=6840CED6",
			link: "",
		},
	],
	中山北路: [
		{ url: "https://placehold.co/600x400?text=ZhongShan+1", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+2", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+3", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+4", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+5", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+6", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+7", link: "" },
		{ url: "https://placehold.co/600x400?text=ZhongShan+8", link: "" },
	],
	民族濱江汽車: [
		{ url: "https://placehold.co/600x400?text=民族濱江汽車+1", link: "" },
		{ url: "https://placehold.co/600x400?text=民族濱江汽車+2", link: "" },
		{ url: "https://placehold.co/600x400?text=民族濱江汽車+3", link: "" },
	],
	// 新北
	新北市商圈聯合發展協會: [
		{
			url: "https://placehold.co/600x400?text=新北市商圈聯合發展協會+1",
			link: "",
		},
		{
			url: "https://placehold.co/600x400?text=新北市商圈聯合發展協會+2",
			link: "",
		},
		{
			url: "https://placehold.co/600x400?text=新北市商圈聯合發展協會+3",
			link: "",
		},
		{
			url: "https://placehold.co/600x400?text=新北市商圈聯合發展協會+4",
			link: "",
		},
	],
	深坑老街商圈: [
		{ url: "https://placehold.co/600x400?text=深坑老街商圈+1", link: "" },
		{ url: "https://placehold.co/600x400?text=深坑老街商圈+2", link: "" },
		{ url: "https://placehold.co/600x400?text=深坑老街商圈+3", link: "" },
		{ url: "https://placehold.co/600x400?text=深坑老街商圈+4", link: "" },
		{ url: "https://placehold.co/600x400?text=深坑老街商圈+5", link: "" },
	],
	鶯歌商圈: [
		{ url: "https://placehold.co/600x400?text=鶯歌商圈+1", link: "" },
		{ url: "https://placehold.co/600x400?text=鶯歌商圈+2", link: "" },
		{ url: "https://placehold.co/600x400?text=鶯歌商圈+3", link: "" },
	],
	板橋府中商圈: [
		{
			url: "https://wefuzhong.com.tw/upload/images/2022_11_28_215914.jpg",
			link: "",
		},
		{
			url: "https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fonelife.tw%2Fimg%2Fa%2F19327b.jpg&f=1&nofb=1&ipt=8cd938ca1da16dc51d7c0ffc13cc35108b9440748e165a2811191a4565f96932",
			link: "",
		},
		{
			url: "https://cpok.tw/wp-content/uploads/2024/09/2024-5.png",
			link: "",
		},
		{
			url: "https://cpok.tw/wp-content/uploads/2023/09/2023-63.jpg",
			link: "",
		},
	],
});

const districts = ref(districtsMap[selectedCity.value]);
const markets = ref([]);

watch(selectedCity, (val) => {
	districts.value = districtsMap[val];
	selectedDistrict.value = districts.value[0];
});
watch(selectedDistrict, (val) => {
	markets.value = marketsMap[val] || [];
	selectedMarket.value = markets.value[0] || "";
});

// 刪除圖片功能
function deleteAd(idx) {
	if (!marketImages.value[selectedMarket.value]) return;
	marketImages.value[selectedMarket.value].splice(idx, 1);
	dialogStore.showNotification("success", "圖片已刪除");
}

// 初始化
selectedDistrict.value = districts.value[0];
markets.value = marketsMap[selectedDistrict.value] || [];
selectedMarket.value = markets.value[0] || "";
</script>

<style scoped lang="scss">
span {
	font-size: 1.5rem;
	color: var(--color-complement-text);
	font-family: var(--font-icon);
	user-select: none;
}
.floating-options-toggle {
	z-index: 2;
	position: absolute;
	top: 0rem;
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
}
.floating-btn {
	background: #fff;
	border: none;
	border-radius: 50%;
	box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
	cursor: pointer;
	transition: box-shadow 0.18s;
	display: flex;
	align-items: center;
	justify-content: center;
	position: relative;
}
.title {
	color: var(--color-normal-text);
	text-align: center;
	justify-content: center;
	display: flex;
	align-items: center;
	margin-block: 1rem;
	margin-top: 2rem;
}
.floating-btn:focus,
.floating-btn:hover {
	box-shadow: 0 4px 18px #2563eb33;
}
.floating-dropdown {
	position: relative;
	background: #fff;
	border-radius: 14px;
	box-shadow: 0 4px 24px rgba(0, 0, 0, 0.13);
	padding-block: 0.4rem;
	padding-inline: 0.6rem;
	display: flex;
	flex-direction: row;
	gap: 0.7rem;
	min-width: 180px;
	animation: fadeIn 0.18s;
}
.floating-dropdown select {
	border-radius: 10px;
	border: 1.5px solid #e0e0e0;
	background: #f8fafc;
	font-size: 1.08rem;
	box-shadow: 0 1px 6px rgba(0, 0, 0, 0.06);
	outline: none;
	transition: border 0.2s, box-shadow 0.2s;
	color: #222;
	font-weight: 500;
}
.floating-dropdown select:focus {
	border: 1.5px solid #2563eb;
	box-shadow: 0 0 0 2px #2563eb22;
}
.floating-dropdown select:hover {
	border: 1.5px solid #60a5fa;
}
@keyframes fadeIn {
	from {
		opacity: 0;
		transform: translateY(-10px);
	}
	to {
		opacity: 1;
		transform: translateY(0);
	}
}
.image-text-view {
	display: flex;
	flex-direction: column;
	align-items: center;
	position: relative;
	height: 100%;
	min-height: 0;
	flex: 1 1 0;
	overflow: auto;
}
.image-section {
	display: flex;
	flex-direction: column;
	align-items: center;
	height: 100%;
	min-height: 0;
	flex: 1 1 0;
	overflow: hidden;
}
.scrollable-images {
	overflow-y: auto;
	overflow-x: hidden;
	display: flex;
	flex-direction: column;
	gap: 1.2rem;
	align-items: center;
}

.image-row {
	position: relative;
}

.image-row img {
	width: 100%;
	border-radius: 12px;
	box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
	object-fit: cover;
}

.delete-btn {
	position: absolute;
	top: 0.5rem;
	right: 0.5rem;
	background: rgba(255, 255, 255, 0.8);
	border-radius: 50%;
	padding: 0.3rem;
	cursor: pointer;
	transition: background 0.2s, transform 0.2s;
	display: flex;
	align-items: center;
	justify-content: center;

	:hover {
		color: var(--color-highlight);
		transform: scale(1.1);
	}
}

.add-btn {
	padding: 0.4rem;
	cursor: pointer;
	display: flex;
	align-items: center;
	justify-content: center;

	&:hover span {
		color: var(--color-highlight);
	}
}

.dialog-content {
	width: 300px;
	display: flex;
	flex-direction: column;

	label {
		margin: 8px 0 4px;
		font-size: var(--font-s);
		color: var(--color-complement-text);
	}
	&-control {
		height: 27px;
		display: flex;
		justify-content: flex-end;
		margin-top: var(--font-ms);

		&-confirm {
			margin: 0 2px;
			padding: 4px 10px;
			border-radius: 5px;
			background-color: var(--color-highlight);
			transition: opacity 0.2s;

			&:hover {
				opacity: 0.8;
			}
		}
	}
}

@media (max-width: 900px) {
	.floating-options {
		top: 1rem;
		right: 1rem;
		padding: 0.3rem 0.5rem;
		gap: 0.5rem;
	}
	.floating-options select {
		font-size: 0.98rem;
		padding: 0.3rem 0.8rem 0.3rem 0.6rem;
	}
}
</style>
