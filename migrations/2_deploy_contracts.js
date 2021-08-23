const PBridge = artifacts.require("PBridge");

module.exports = async function (deployer) {

  await deployer.deploy(PBridge, ['0xB74cA291733f155e39AaCfa6B217ceB41a2204c2', '0x250C60ba4362a20f2018c1552E485d0384003627', '0x12bbcD0C5F548F1FF2b29ACB67Fac74860c6980e']);

}
