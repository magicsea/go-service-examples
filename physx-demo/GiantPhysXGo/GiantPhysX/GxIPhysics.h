/*
 * This file is part of the GiantPhysX package.
 *
 * Copyright (C) 2019, GIANT - liangjinyao@ztgame.com. ALL RIGHTS RESERVED.
 */
#pragma once

#include "GiantPhysX/GxAPIForward.h"

namespace GiantPhysX
{
	/// <summary>
	/// ����ϵͳ��
	/// �ṩ����ϵͳ��ȫ�����Լ�����
	/// </summary>
	class GX_API GxIPhysics
	{
	public:
		/// <summary>
		/// ��������
		/// </summary>
		virtual ~GxIPhysics() {}

		/// <summary>
		/// �����³���
		/// </summary>
		/// <param name="config">���������ļ�·��</param>
		/// <returns>�³�������</returns>
		virtual GxIScene* CreateScene(const char* config = nullptr) = 0;

		/// <summary>
		/// ���ٳ���
		/// </summary>
		/// <remarks>
		/// ��������ʱ���ó��������ж��󶼻ᱻ���١�
		/// </remarks>
		/// <param name="scene">Ҫ���ٵĳ�������</param>
		virtual void DestroyScene(GxIScene* scene) = 0;
	};

} // namespace GiantPhysX