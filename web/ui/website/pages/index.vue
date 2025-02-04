<template>
	<div class="page">
		<header>
			<Header></Header>

			<section class="hero-section">
				<div class="hero-section--cta">
					<h1>A Cloud native Webhook Service</h1>
					<p>With out-of-the-box security, reliability and scalability for your webhooks infrastructure.</p>
					<nuxt-link to="/docs">Get Started</nuxt-link>

					<form @submit.prevent="requestAccess()">
						<div class="input">
							<input type="email" id="email" placeholder="Enter your email" aria-label="Email" v-model="earlyAccessEmail" />
						</div>
						<button :disabled="isSubmitingloadingEarlyAccessForm">{{ isSubmitingloadingEarlyAccessForm ? '...requesting' : earlyAccessFormButtonText }}</button>
					</form>
				</div>
				<div class="hero-section--img">
					<img src="~/assets/images/hero-img.png" alt="hero" />
				</div>
			</section>
		</header>

		<section class="section companies">
			<p>Companies using us:</p>
			<ul>
				<li>
					<a target="_blank" rel="noopener noreferrer" href="https://buycoins.africa/">
						<img src="~/assets/images/buycoins-logo.svg" style="width: 70px" alt="buycoins logo" />
					</a>
				</li>
				<li>
					<a target="_blank" rel="noopener noreferrer" href="https://termii.com/">
						<img src="~/assets/images/termii-logo.png" style="width: 80px" alt="termii logo" />
					</a>
				</li>
				<li>
					<a target="_blank" rel="noopener noreferrer" href="https://www.getwallets.co/">
						<img src="~/assets/images/getwallets-logo.svg" style="width: 80px" alt="getwallets logo" />
					</a>
				</li>
				<li>
					<a target="_blank" rel="noopener noreferrer" href="https://www.dojah.io/">
						<img src="~/assets/images/dojah-logo.png" style="width: 50px" alt="dojah logo" />
					</a>
				</li>
			</ul>
		</section>

		<section class="section features" id="features">
			<div class="container">
				<h2 class="section--title">Why Convoy</h2>
				<p class="section--description">
					Convoy enables you get started with delivering webhooks products in the space of minutes, along with a number of carefully thought out features to truly enable/enhance your webhooks delivery
					at scale.
				</p>

				<ul>
					<li>
						<img src="~/assets/images/secure-feature-icon.svg" alt="secure feature icon" />
						<h4>Secure</h4>
						<p>Create secrets, sign payload, verify event. Increase security by enabling rolling secrets.</p>
					</li>
					<li>
						<img src="~/assets/images/ease-feature-icon.svg" alt="ease feature icon" />
						<h4>Ease of use</h4>
						<p>Easily filter & debug events sent to multiple applications & endpoints with Delivery Attempt Logs</p>
					</li>
					<li>
						<img src="~/assets/images/cloud-feature-icon.svg" alt="cloud feature icon" />
						<h4>Cloud & Language Agnostic</h4>
						<p>Deploy to any environment and send events from any language (Ruby, Golang, JavaScript etc)</p>
					</li>
					<li>
						<img src="~/assets/images/rate-feature-icon.svg" alt="rate feature icon" />
						<h4>Rate Limiting</h4>
						<p>Avoid blasting too much events to a single endpoint at once with flexible rate limiting controls</p>
					</li>
					<li>
						<img src="~/assets/images/scalable-feature-icon.svg" alt="scalable feature icon" />
						<h4>Scalable</h4>
						<p>Independently scale Convoy as your system needs grows.</p>
					</li>
					<li>
						<div class="features--dashboard-screenshot">
							<img src="~/assets/images/dashboard-img.png" alt="dashboard image" />
						</div>
					</li>
				</ul>
				<div class="features--dashboard-screenshot stand-alone">
					<img src="~/assets/images/dashboard-img.png" alt="dashboard image" />
				</div>
			</div>
		</section>

		<section class="section core-features">
			<div class="container">
				<div>
					<h2 class="section--title">Core Features</h2>
					<ul>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Independently scalable</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Signatures & Rolling secrets</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Retries</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Clustering</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Language Agnostic</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Idempotency</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>URL per Event Type</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Static IPs</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Rich UI - Event Logs & Querying.</p>
						</li>
						<li>
							<img src="~/assets/images/feature-check-icon.svg" alt="check icon" />
							<p>Payload Filtering</p>
						</li>
					</ul>
				</div>
				<div>
					<div class="core-features--img-container">
						<img src="~/assets/images/core-features-img.svg" alt="core flow" />
					</div>
				</div>
			</div>
		</section>

		<section class="section play-ground">
			<div class="container">
				<h2>Convoy Play Ground</h2>
				<p>We've deployed an instance of Convoy to the cloud so you can play with it and have a taste of how Convoy will enable your webhook service</p>
				<div class="img">
					<a href="https://playground.getconvoy.io/" target="_blank">
						Try it out free
						<img src="~/assets/images/arrow-right-icon.svg" alt="arrow right" />
					</a>
					<img src="~/assets/images/playground-img.png" alt="dashboard screenshot" />
				</div>
			</div>
			<div class="bg"></div>
		</section>

		<footer>
			<div class="container">
				<nav>
					<div class="logo">
						<img src="~/assets/images/logo.svg" alt="logo" />
						<p>
							by
							<a href="https://frain.dev">Frain</a>
						</p>
					</div>

					<ul class="socials">
						<li>
							<a target="_blank" rel="noopener noreferrer" href="https://join.slack.com/t/convoy-community/shared_invite/zt-xiuuoj0m-yPp~ylfYMCV9s038QL0IUQ">
								<img src="~/assets/images/slack-icon.svg" alt="slack logo" />
							</a>
						</li>
						<li>
							<a target="_blank" rel="noopener noreferrer" href="https://twitter.com/fraindev"><img src="~/assets/images/twitter-icon.svg" alt="twitter logo" /></a>
						</li>
						<li>
							<a target="_blank" rel="noopener noreferrer" href="https://github.com/frain-dev/convoy"><img src="~/assets/images/github-icon.svg" alt="mail logo" /></a>
						</li>
						<li>
							<a target="_blank" rel="noopener noreferrer" href="mailto:info@frain.dev"><img src="~/assets/images/mail-icon.svg" alt="mail logo" /></a>
						</li>
					</ul>
				</nav>
				<p>Copyright 2022, All Rights Reserved</p>
			</div>
		</footer>
	</div>
</template>

<script>
export default {
	data() {
		return {
			isSubmitingloadingEarlyAccessForm: false,
			earlyAccessFormButtonText: 'Convoy Cloud Early Access',
			earlyAccessEmail: ''
		};
	},
	async asyncData({ $content, params }) {
		const pageData = await $content('docs').fetch();
		return { pageData };
	},
	methods: {
		async requestAccess() {
			this.isSubmitingloadingEarlyAccessForm = true;
			try {
				const response = await fetch('/.netlify/functions/subscribe', {
					method: 'POST',
					mode: 'cors',
					cache: 'no-cache',
					credentials: 'same-origin',
					headers: {
						'Content-Type': 'application/json'
					},
					redirect: 'follow',
					referrerPolicy: 'no-referrer',
					body: JSON.stringify({
						email: this.earlyAccessEmail
					})
				});
				await response.json();
				this.earlyAccessFormButtonText = 'Request Sent';
				this.setDefaultAccessButtonText();
				this.isSubmitingloadingEarlyAccessForm = false;
			} catch (error) {
				this.earlyAccessFormButtonText = 'Error';
				this.setDefaultAccessButtonText();
				this.isSubmitingloadingEarlyAccessForm = false;
			}
		},
		setDefaultAccessButtonText() {
			setTimeout(() => {
				this.earlyAccessFormButtonText = 'Convoy Cloud Early Access';
			}, 5000);
		}
	}
};
</script>

<style lang="scss" scoped>
$desktopBreakPoint: 880px;

header {
	background: url('~/assets/images/header-bg-pattern.svg') #1e252b no-repeat;
	background-size: cover;
	padding: 32px 20px 0;

	.hero-section {
		max-width: 1106px;
		margin: auto;
		display: flex;
		justify-content: space-between;
		min-height: 603px;
		height: 100%;
		flex-direction: column;
		padding: 75px 0 100px;

		@media (min-width: $desktopBreakPoint) {
			flex-direction: row;
			align-items: center;
		}

		&--cta {
			max-width: 533px;

			h1 {
				font-weight: bold;
				font-size: 48px;
				line-height: 64px;
				color: #ffffff;
				margin-bottom: 16px;
			}

			p {
				font-size: 18px;
				line-height: 30px;
				color: #ffffff;
			}
		}

		&--img {
			margin-left: 50px;
			text-align: center;
			display: none;

			img {
				max-width: 330px;
			}

			@media (min-width: $desktopBreakPoint) {
				display: block;
			}
		}

		button,
		a {
			background: #477db3;
			padding: 9px 21px;
			color: #ffffff;
			font-weight: 500;
			font-size: 14px;
			line-height: 28px;
			margin: 20px 0 0;
			max-width: 200px;
			width: 100%;
			display: block;
			border-radius: 5px;
			text-align: center;
		}

		form {
			background: #ffffff;
			max-width: 430px;
			display: flex;
			border-radius: 8px;
			margin-top: 40px;
			flex-direction: column;

			@media (min-width: $desktopBreakPoint) {
				flex-direction: row;
			}

			.input {
				width: 100%;
				display: flex;

				input {
					max-width: 278px;
					width: 100%;
					border: none;
					padding: 9px 23px;
					border-radius: 8px 0 0 8px;
					outline: none;
					margin: 0;

					&::placeholder {
						font-weight: 500;
						font-size: 14px;
						line-height: 28px;
						color: #bdbdbd;
					}
				}
			}

			button {
				max-width: none;
				width: 100%;
				margin: 0;
				border-radius: 0 0 8px 8px;
				margin-top: 10px;
				font-size: 12px;

				@media (min-width: $desktopBreakPoint) {
					max-width: 220px;
					width: 100%;
					margin-top: 0;
					border-radius: 0 8px 8px 0;
				}
			}
		}
	}
}

.section {
	.container {
		max-width: 1150px;
		margin: auto;
	}

	.section--title {
		font-weight: bold;
		font-size: 48px;
		line-height: 60px;
		letter-spacing: -1.584px;

		@media (max-width: 1089px) {
			font-size: 40px;
		}
	}

	.section--description {
		font-size: 18px;
		line-height: 30px;
	}
}

section.companies {
	display: flex;
	max-width: 1150px;
	margin: -30px auto 0;
	justify-content: space-between;
	padding: 20px 20px 25px;
	background: #fafafe;
	border-radius: 8px 8px 0px 0px;
	align-items: center;
	flex-wrap: wrap;

	p {
		font-weight: bold;
		font-size: 16px;
		line-height: 24px;
		padding: 0;
		margin: 0 0 20px;
	}

	ul {
		display: flex;
		align-items: center;
		margin-bottom: 10px;
		flex-wrap: wrap;

		li {
			&:not(:last-of-type) {
				margin-right: 50px;
			}

			img {
				filter: grayscale(100%);
				transition: all 0.3s ease-in-out;
				margin-bottom: 20px;

				&:hover {
					filter: grayscale(0%);
				}
			}
		}
	}
}

.features {
	padding: 160px 20px 800px;

	@media (max-width: 1089px) {
		padding: 60px 20px 100px;
	}

	.container {
		position: relative;
	}

	.section--title {
		margin-bottom: 26px;
	}

	.section--description {
		max-width: 608px;
		margin-bottom: 77px;
	}

	ul {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(238px, 1fr));
		gap: 80px 32px;
		margin-bottom: 50px;

		img {
			width: 70px;
			margin-bottom: 36px;
		}

		li {
			position: relative;

			@media (max-width: 1089px) {
				.features--dashboard-screenshot {
					display: none;
				}
			}
		}

		h4 {
			font-weight: bold;
			font-size: 18px;
			line-height: 22px;
			letter-spacing: -0.5px;
			margin-bottom: 16px;
		}

		p {
			font-size: 15px;
			line-height: 22px;

			@media (max-width: 1089px) {
				font-size: 16px;
			}
		}
	}

	.features--dashboard-screenshot {
		position: absolute;
		left: 0;
		top: 0;
		width: 1100px;
		height: fit-content;

		@media (max-width: 1644px) {
			width: 1000px;
		}

		@media (max-width: 1443px) {
			width: 800px;
		}

		img {
			width: 100%;
		}

		&.stand-alone {
			position: unset;
			width: 100%;
			text-align: center;
			display: none;

			@media (max-width: 1089px) {
				display: block;
			}

			img {
				width: 100%;
			}
		}
	}
}

.core-features {
	background: #2b4b6a;
	padding: 80px 20px;
	color: #ffffff;
	margin-bottom: 490px;

	@media (min-width: $desktopBreakPoint) {
		margin-bottom: 590px;
	}

	.container {
		display: flex;
		align-items: center;
		justify-content: space-between;

		@media (max-width: 1089px) {
			display: unset;
		}
	}

	ul {
		margin-top: 32px;

		li {
			display: flex;
			align-items: center;
			margin-bottom: 29px;
			position: relative;
			font-weight: 500;
			font-size: 18px;
			line-height: 27px;

			&:not(:last-of-type)::after {
				content: '';
				position: absolute;
				height: 20px;
				width: 1px;
				background: #4b6680;
				bottom: -24px;
				left: 11px;
			}
		}

		img {
			width: 24px;
			margin-right: 24px;
		}
	}

	&--img-container {
		background: linear-gradient(0deg, #ffffff, #ffffff);
		box-shadow: 40px 44px 81px rgba(111, 118, 138, 0.160784);
		border-radius: 16px;
		height: 950px;
		max-width: 649px;
		width: 100%;
		margin-top: -160px;
		margin-bottom: -160px;
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 0 40px;

		@media (max-width: 1089px) {
			margin: 100px auto 0;
			height: fit-content;
			padding: 40px 20px;
		}

		img {
			max-width: 467px;
			width: 100%;
		}
	}
}

.play-ground {
	padding-bottom: 300px;
	position: relative;

	@media (min-width: $desktopBreakPoint) {
		padding-bottom: 400px;
	}

	.bg {
		content: '';
		bottom: 0;
		left: -300px;
		right: -300px;
		height: 300px;
		background: #ffffff;
		border-radius: 100% 100% 0 0;
		margin-bottom: -400px;
		margin-right: -300px;
		margin-left: -300px;

		@media (min-width: $desktopBreakPoint) {
			height: 460px;
		}
	}

	.container {
		width: 90%;
		box-shadow: 0px 15px 50px rgba(21, 37, 72, 0.05);
		border-radius: 24px;
		text-align: center;
		padding: 80px 2rem 0;
		background: url('~/assets/images/playground-bg.svg') #ffffff no-repeat;
		background-size: contain;
		background-position: center;
		position: absolute;
		left: 50%;
		top: -300px;
		transform: translate(-50%, 0);

		@media (min-width: $desktopBreakPoint) {
			background-position: top center;
		}
	}

	h2 {
		font-weight: bold;
		font-size: 34px;
		line-height: 60px;

		@media (min-width: $desktopBreakPoint) {
			font-size: 48px;
		}
	}

	p {
		font-weight: 500;
		font-size: 18px;
		line-height: 30px;
		max-width: 620px;
		margin: 16px auto 26px;
	}

	.img {
		max-width: 830px;
		margin: auto;
		position: relative;

		img {
			width: 100%;
			margin-bottom: -4px;
		}

		a {
			position: absolute;
			left: 50%;
			top: 64%;
			transform: translate(-50%, -50%);
			padding: 17px 24px;
			background: #477db3;
			color: #ffffff;
			width: 212px;
			text-align: left;
			border-radius: 8px;
			display: flex;
			align-items: center;
			justify-content: space-between;
			font-weight: bold;
			font-size: 18px;
			line-height: 22px;

			&::before {
				content: '';
				position: absolute;
				width: 86px;
				height: 102px;
				right: -40px;
				background: url('~/assets/images/play-ground-button-bg.svg');
				background-size: contain;
				z-index: -1;
			}

			img {
				width: 12px;
			}
		}
	}
}
</style>
